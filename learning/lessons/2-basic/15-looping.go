package main

import (
	"fmt"
)

func main() {

	fmt.Println("###looping collections###")

	fmt.Println("#loop fors#")
	sl := []int{10, 20, 30, 40}
	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}

	fmt.Println("#ranging#")
	for i, value := range sl {
		fmt.Println(i, value)
	}

	fmt.Println("#loop maps#")
	prodPrice := map[string]int{
		"widget":             75,
		"turbo widget":       100,
		"convertible widget": 150,
	}
	for key := range prodPrice {
		fmt.Println(key)
	}
	for _, value := range prodPrice {
		fmt.Println(value)
	}
	for key, value := range prodPrice {
		fmt.Println(key, value)
	}
}
