package lemon_promise

type SimplePromise struct {
	result any
}

func NewSimplePromise(result any) *SimplePromise {
	return &SimplePromise{result: result}
}

func (s *SimplePromise) Then(callback PipeDoneCallback) Promise {
	return NewPipePromise(s, callback)
}

func (s *SimplePromise) Done(callback DoneCallback) Promise {
	callback(s.result)
	return s
}

func (s *SimplePromise) Fail(callback FailCallBack) Promise {
	return s
}

func (s *SimplePromise) Get() any {
	return s.result
}
