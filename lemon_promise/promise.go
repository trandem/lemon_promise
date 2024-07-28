package lemon_promise

type PipeDoneCallback func(r any) Promise
type DoneCallback func(r any) any
type FailCallBack func(e error) Promise

type Result struct {
	result any
}

type Promise interface {
	Then(callback PipeDoneCallback) Promise
	Done(callback DoneCallback) Promise
	Get() any
}

type SimpleDonePromise struct {
	result any
	err    error
}

func (s *SimpleDonePromise) Done(callback DoneCallback) Promise {
	callback(s.result)
	return s
}

func (s *SimpleDonePromise) Get() any {
	return s.result
}

func NewSimpleDonePromise(result any) *SimpleDonePromise {
	return &SimpleDonePromise{result: result}
}

func (s *SimpleDonePromise) Then(callback PipeDoneCallback) Promise {
	out := callback(s.result)
	return NewSimpleDonePromise(out)
}

type CompletePromise struct {
	result           any
	pipeDoneCallBack PipeDoneCallback
	doneCallback     DoneCallback
	nextPromise      *CompletePromise
}

func (c *CompletePromise) Done(callback DoneCallback) Promise {
	if c.IsDone() {
		callback(c.result)
	}
	c.doneCallback = callback
	return c
}

func (c *CompletePromise) Get() any {
	return c.result
}

func (c *CompletePromise) PromiseDone(rs any) {
	c.result = rs
	if c.doneCallback != nil {
		c.doneCallback(c.result)
	}

	if c.pipeDoneCallBack != nil {
		c.pipeDoneCallBack(c.result).Done(func(r any) any {
			c.nextPromise.PromiseDone(r)
			return nil
		})
	}
}

func (c *CompletePromise) IsDone() bool {
	return c.result != nil
}
func (c *CompletePromise) Then(callback PipeDoneCallback) Promise {
	nextPromise := &CompletePromise{}
	c.pipeDoneCallBack = callback
	c.nextPromise = nextPromise
	if c.IsDone() {
		callback(c.result).Done(func(r any) any {
			nextPromise.result = r
			return nil
		})
	}
	return nextPromise
}
