package main

import (
	"errors"
	"fmt"
	"log"
)

func volume(r float64) (float64, error) {

	if r < 0 {
		//func New(text string) error
		return 0, errors.New("Volume calculation failed; radius negative")
	}
	return (4.0 / 3.0) * 3.14 * r * r * r, nil
}

func main() {

	radius := -1.0
	vol, err := volume(radius)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("volume of sphere is %0.2f", vol)
}
