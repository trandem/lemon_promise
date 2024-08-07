package main

import (
	"errors"
	"fmt"
	"lemon_promise/lemon_promise"
)

func main() {
	c := lemon_promise.NewCompletablePromise()

	lemon_promise.NewSimplePromise("444").
		Then(func(r any) lemon_promise.Promise {
			return c
		}).
		Then(func(r any) lemon_promise.Promise {
			return lemon_promise.NewSimplePromise(r)
		}).
		Done(func(r any) {
			fmt.Println("done")
			fmt.Println(r)
		}).
		Fail(func(e error) {
			fmt.Println("fail")
			fmt.Println(e)
		})

	c.Reject(errors.New("tet"))
}
