package lemon_promise

import "sync"

type CompletablePromise struct {
	result       any
	err          error
	doneCallBack DoneCallback
	failCallBack FailCallBack
	mutex        *sync.Mutex
}

func NewCompletablePromise() *CompletablePromise {
	mutex := &sync.Mutex{}
	return &CompletablePromise{mutex: mutex, result: nil,
		failCallBack: nil, doneCallBack: nil}
}

func (c *CompletablePromise) Resolve(result any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.result = result
	if c.doneCallBack != nil {
		c.doneCallBack(c.result)
	}
}

func (c *CompletablePromise) Reject(err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.err = err
	if c.failCallBack != nil {
		c.failCallBack(err)
	}
}

func (c *CompletablePromise) Then(callback PipeDoneCallback) Promise {
	return NewPipePromise(c, callback)
}

func (c *CompletablePromise) Done(callback DoneCallback) Promise {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.doneCallBack = callback
	if c.result != nil {
		c.doneCallBack(c.result)
	}
	return c
}

func (c *CompletablePromise) Fail(callback FailCallBack) Promise {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.failCallBack = callback
	if c.err != nil {
		c.failCallBack(c.err)
	}
	return c
}

func (c *CompletablePromise) Get() any {
	return c.result
}
