package lemon_promise

import "sync"

type PipePromise struct {
	result       any
	err          error
	doneCallBack DoneCallback
	failCallBack FailCallBack
	mutex        *sync.Mutex

	pipeDoneCallBack PipeDoneCallback
	previousPromise  Promise
}

func NewPipePromise(promise Promise, pipeDone PipeDoneCallback) *PipePromise {
	mutex := &sync.Mutex{}

	p := &PipePromise{previousPromise: promise, mutex: mutex}

	promise.Done(func(r any) {
		if pipeDone != nil {
			pipeDone(r).Done(func(r any) {
				p.Resolve(r)
			}).Fail(func(e error) {
				p.Reject(e)
			})
		}
	}).Fail(func(e error) {
		p.Reject(e)
	})

	return p
}

func (p *PipePromise) Resolve(result any) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.result = result
	if p.doneCallBack != nil {
		p.doneCallBack(p.result)
	}
}

func (p *PipePromise) Reject(err error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.err = err
	if p.failCallBack != nil {
		p.failCallBack(err)
	}
}

func (p *PipePromise) Then(callback PipeDoneCallback) Promise {
	return NewPipePromise(p, callback)
}

func (p *PipePromise) Done(callback DoneCallback) Promise {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.doneCallBack = callback
	if p.result != nil {
		p.doneCallBack(p.result)
	}
	return p
}

func (p *PipePromise) Fail(callback FailCallBack) Promise {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.failCallBack = callback
	if p.err != nil {
		p.failCallBack(p.err)
	}
	return p
}

func (p *PipePromise) Get() any {
	return p.result
}
