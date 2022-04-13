package main

import (
	"fmt"
)

type Shape interface {
	volume() float64
}

type Cube struct {
	depth  float64
	width  float64
	height float64
}

func (cu *Cube) volume() float64 {
	return cu.depth * cu.width * cu.height
}

type Sphere struct {
	radius float64
}

func (sp *Sphere) volume() float64 {
	return ((4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius))
}

func totalVolume(shapes ...Shape) float64 {
	var volume float64
	for _, s := range shapes {
		volume += s.volume()
	}
	return volume
}

func main() {
	fmt.Println("###managing objects###")

	fmt.Println("#operate#")
	cu := Cube{depth: 4, width: 4, height: 4}
	sp := Sphere{radius: 2}
	fmt.Println("Total Volume: ", totalVolume(&cu, &sp))
}
