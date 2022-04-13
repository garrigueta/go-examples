package main

import (
	"fmt"
)

func main() {

	fmt.Println("###program flow###")
	fmt.Println("#def & init#")
	// TODO: (useless??)
	if score := 99; score < 100 {
		fmt.Println("oops!")
	} else if score >= 100 && score < 1000 {
		fmt.Println("good!")
	} else {
		fmt.Println("wow")
	}

	fmt.Println("#switch#")
	workday := 3
	switch workday {
	case 1:
		fmt.Println("day1")
	case 2:
		fmt.Println("day2")
	case 3:
		fmt.Println("day3")
	case 4:
		fmt.Println("day4")
	case 5:
		fmt.Println("day5")
	default:
		fmt.Println("alldays")
	}

	fmt.Println("#switch true#")
	score := 99
	switch {
	case score < 100:
		fmt.Println("oops!")
	case score >= 100 && workday < 100:
		fmt.Println("good!")
	default:
		fmt.Println("wow!")
	}
}
