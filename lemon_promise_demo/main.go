package main

import "fmt"

func main() {
	p1 := NewSimpleDonePromiseDemo("11111")
	p1.Done(func(r any) any {
		fmt.Printf("%+v\n", 13232)
		return NewSimpleDonePromiseDemo(11111)
	})

	fmt.Printf("%+v\n", p1.Get())
}
