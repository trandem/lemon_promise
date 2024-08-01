package main

type SimpleDonePromiseDemo struct {
	value any
}

func NewSimpleDonePromiseDemo(value any) *SimpleDonePromiseDemo {
	return &SimpleDonePromiseDemo{value: value}
}

func (promise *SimpleDonePromiseDemo) Done(callback DoneCallbackDemo) PromiseDemo {
	callback(promise.value)
	return promise
}

func (promise *SimpleDonePromiseDemo) Fail(callback FailCallbackDemo) PromiseDemo {
	// SimpleDonePromise ignore case Fail
	return promise
}

func (promise *SimpleDonePromiseDemo) Get() any {
	return promise.value
}
