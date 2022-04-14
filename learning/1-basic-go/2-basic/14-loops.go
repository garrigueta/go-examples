package main

import (
	"fmt"
)

func main() {

	fmt.Println("###loops & branches###")

	fmt.Println("#count#")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("#same#")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("#break#")
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("#continue#")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("#outerlabel#")
	i = 0
outerlabel:
	for i < 5 {
		if i == 2 {
			// uses goto statement
			i++
			goto outerlabel
		}
		fmt.Println(i)
		i++
	}

	fmt.Println("#infinite#")
	i = 0
	for {
		if i == 3 {
			break
		}
		fmt.Println(i)
		i++
	}
}
