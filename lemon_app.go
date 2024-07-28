package main

import (
	"fmt"
	"lemon_promise/lemon_promise"
)

func main() {
	cP := &lemon_promise.CompletePromise{}
	x1 := &lemon_promise.CompletePromise{}

	x := cP.Then(func(r any) lemon_promise.Promise {
		return x1
	})

	x2 := x.Then(func(r any) lemon_promise.Promise {
		fmt.Printf("lolol %+v\n", r)
		return lemon_promise.NewSimpleDonePromise(r)
	})

	cP.PromiseDone("abc")
	fmt.Printf("%+v\n", cP.Get())
	fmt.Printf("%+v\n", x.Get())
	x1.PromiseDone("444")
	fmt.Printf("%+v\n", x.Get())
	fmt.Printf("%+v\n", x2.Get())
	// https://github.com/dungba88/promise4j
	// survey full flow and write flow code
}
