package main

import (
	f "fmt"
	"math"
	. "os"
	_ "strings"

	m "example.local/calc/math"
)

func main() {
	f.Println(m.Add(10, 20))
	f.Println(m.Pi)
	f.Println(math.Pi)
	f.Fprintln(Stdout, "Hello")
}
