package lemon_promise

import "sync"

type JoinPromise struct {
	result       any
	err          error
	doneCallBack DoneCallback
	failCallBack FailCallBack
	mutex        *sync.Mutex
}

type JoinResult struct {
	Results []any
}

func NewJoinPromise(promises ...Promise) *JoinPromise {
	mutex := &sync.Mutex{}
	listResult := make([]any, len(promises))

	jp := JoinPromise{mutex: mutex}
	c := 0
	for i, p := range promises {
		p.Done(func(r any) {
			mutex.Lock()
			listResult[i] = r
			c++
			mutex.Unlock()
			if c == len(promises) {
				jp.Resolve(&JoinResult{Results: listResult})
			}
		}).Fail(func(e error) {
			jp.Reject(e)
		})
	}

	return &jp
}

func (j *JoinPromise) Resolve(result any) {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.result = result
	if j.doneCallBack != nil {
		j.doneCallBack(j.result)
	}
}

func (j *JoinPromise) Reject(err error) {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.err = err
	if j.failCallBack != nil {
		j.failCallBack(err)
	}
}

func (j *JoinPromise) Then(callback PipeDoneCallback) Promise {
	return NewPipePromise(j, callback)
}

func (j *JoinPromise) Done(callback DoneCallback) Promise {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.doneCallBack = callback
	if j.result != nil {
		j.doneCallBack(j.result)
	}
	return j
}

func (j *JoinPromise) Fail(callback FailCallBack) Promise {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	j.failCallBack = callback
	if j.err != nil {
		j.failCallBack(j.err)
	}
	return j
}

func (j *JoinPromise) Get() any {
	return j.result
}
