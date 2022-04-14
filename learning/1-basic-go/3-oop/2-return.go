package main

import (
	"fmt"
)

func add(x, y int) (a, b int, c bool) {

	a = x + b
	b = x - y
	c = x > y
	//return a, b, false
	return
}

func main() {

	fmt.Println("###return with functions###")

	fmt.Println("#typed returns#")
	x := 20
	y := 10
	fmt.Println(add(x, y))

}
