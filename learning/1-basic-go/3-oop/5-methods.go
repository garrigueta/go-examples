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

// volume method to calculate volume of a Box
func (b *Box) volume() float64 {

	return b.D * b.W * b.H
}

func main() {

	fmt.Println("###managing methods###")

	fmt.Println("#operate#")
	b := Box{D: 5, W: 4, H: 3}
	fmt.Println(b.volume())
}
