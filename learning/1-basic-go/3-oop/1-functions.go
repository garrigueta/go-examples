//DRY: don't repeat yourself
package main

import (
	"fmt"
)

var x int = 20
var y int = 10

func avg(s []float64) float64 {
	total := 0.0

	for _, val := range s {
		total += val
	}

	return total / float64(len(s))

}

func add1() int {

	return x + y

}

func add2(x, y, z int) int {

	return x + y + z

}

func add3(x []int) int {

	total := 0
	for _, val := range x {
		total += val
	}
	return total

}

func multiply(x, y *int) int {

	return *x * *y

}

func main() {

	fmt.Println("###handling functions###")

	fmt.Println("#functions#")
	s := []float64{1.8, 2.2, 2.0}
	fmt.Println(avg(s))

	fmt.Println("#functions with params#")
	x := 20
	y := 10
	z := 10
	fmt.Println(add1())
	fmt.Println(add2(x, z, y))
	slc := []int{10, 20, 5}
	fmt.Println(add3(slc))
	fmt.Println(multiply(&x, &y))

	fmt.Println("#returning data#")

}
