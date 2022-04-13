package main

import (
	"fmt"
)

func add(values ...int) {
	total := 0
	for _, val := range values {
		total += val
	}
	fmt.Println(values, total)
}

func main() {

	fmt.Println("###variadic functions###")

	fmt.Println("#packed operator#")
	vals := []int{1, 2, 3, 4}
	add(vals...)
	vals = []int{1, 2, 3, 4, 5, 6}
	add(vals...)
}
