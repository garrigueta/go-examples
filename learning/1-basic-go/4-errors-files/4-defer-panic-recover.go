package main

import (
	"fmt"
)

//func panic(v interface{})
//func recover() interface{}

func p(s string, i int) {
	//x := [3]int{1, 2, 3}
	//x[i] = 11
	if 1 > 2 {
		panic(s)
	}
}

func p2(s string, i int) {
	x := [3]int{1, 2, 3}
	x[i] = 11
	//if 1 > 2 {
	//	panic(s)
	//}
}

func r() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("Recovered from panic")
	}
}

func main() {
	// Defer is a process that places a function call onto a stack.
	// Each deferred function is executed in reverse order when the
	// calling function completes whether a panic is called or not.
	defer r()
	p("runtime error: enter panic state", 3)
	p2("runtime error: enter panic state", 3)
}
