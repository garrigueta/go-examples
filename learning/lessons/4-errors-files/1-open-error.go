package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//signature: func Open(name string) (file *File, err error)
	f, err := os.Open("myFakeFile.txt")
	defer f.Close()

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("file successfully opened:", f.Name())
}
