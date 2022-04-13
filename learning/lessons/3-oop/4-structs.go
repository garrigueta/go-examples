package main

import (
	"fmt"
)

// Box is a cool object
type Box struct {
	D float64
	W float64
	H float64
}

func main() {

	fmt.Println("###managing structs###")

	fmt.Println("#operate#")
	b := Box{5, 4, 4}
	fmt.Println(b)
	b2 := Box{D: 5, W: 4, H: 4}
	fmt.Println(b2)
	b.D = 6
	fmt.Println(b)
	ptr := &b
	(*ptr).D = 7
	fmt.Println(b)
	ptr.D = 8
	fmt.Println(b)
}
