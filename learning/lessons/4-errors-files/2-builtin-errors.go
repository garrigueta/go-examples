package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	if _, err := os.Open("myFakeFile.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file does not exist")
		} else {
			log.Println(err)
		}
		return
	}

	fmt.Println("file successfully opened")
}
