package main

type PromiseDemo interface {
	Done(callback DoneCallbackDemo) PromiseDemo
	Fail(callback FailCallbackDemo) PromiseDemo
	Get() any
}

type DoneCallbackDemo func(r any) any

type FailCallbackDemo func(r any) any
