package main

import (
	"fmt"
	_ "strings"

	"example.local/calc/math"
)

func main() {
	values := []float64{1, 10, 100, 1000}
	fmt.Println(math.Add(10, 20))
	fmt.Println(math.Avg(values))
}
