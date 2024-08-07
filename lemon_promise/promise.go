package lemon_promise

type PipeDoneCallback func(r any) Promise
type DoneCallback func(r any)
type FailCallBack func(e error)

type Promise interface {
	Then(callback PipeDoneCallback) Promise
	Done(callback DoneCallback) Promise
	Fail(callback FailCallBack) Promise
	Get() any
}

type Deferred interface {
	Resolve(result any)
	Reject(err error)
}
