package main

import (
	"fmt"
)

func main() {
	defer fmt.Print("second")

	fmt.Println("first")
}
