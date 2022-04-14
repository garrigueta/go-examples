package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("###string functions###")

	fmt.Println("#contains#")
	fmt.Println(strings.Contains("Working with string functions", "phunctions"))

	fmt.Println("#replace#")
	fmt.Println(strings.Replace("Working with string functions", "functions", "variables", -1))

	fmt.Println("#title#")
	fmt.Println(strings.Title("Working with string functions"))

	fmt.Println("#trim#")
	fmt.Println(strings.Trim("___Working with string functions___", "_"))

}
