package main

import (
	"fmt"

	"example.local/calc/math"
	"example.local/calc/math/stats"
)

func main() {
	values := []float64{1, 10, 100, 1000}
	fmt.Println(math.Add(10, 20))
	fmt.Println(stats.Avg(values))
}
