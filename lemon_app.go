package main

import (
	"errors"
	"fmt"
	"lemon_promise/lemon_promise"
)

func main() {
	ex1()
	ex2()
	ex3()
}
func ex3() {
	p1 := lemon_promise.NewCompletablePromise()
	p2 := lemon_promise.NewCompletablePromise()
	p3 := lemon_promise.NewSimplePromise("abc")

	lemon_promise.NewJoinPromise(p1, p2, p3).Done(func(r any) {
		joinRs := r.(*lemon_promise.JoinResult)
		listRs := joinRs.Results
		fmt.Println(listRs[0])
		fmt.Println(listRs[1])
		fmt.Println(listRs[2])
	}).
		Fail(func(e error) {
			fmt.Println(e)
		})
	p2.Resolve("2")
	p1.Reject(errors.New("test"))
}

func ex2() {
	p1 := lemon_promise.NewCompletablePromise()
	p2 := lemon_promise.NewCompletablePromise()
	p3 := lemon_promise.NewSimplePromise("abc")

	lemon_promise.NewJoinPromise(p1, p2, p3).Done(func(r any) {
		joinRs := r.(*lemon_promise.JoinResult)
		listRs := joinRs.Results
		fmt.Println(listRs[0])
		fmt.Println(listRs[1])
		fmt.Println(listRs[2])
	})
	p2.Resolve("2")
	p1.Resolve(1)
}

func ex1() {
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
