package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	// formerly ioutil.ReadFile()
	// func ReadFile(name string) ([]byte, error)
	contents, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// type cast slice of uint8 values in "contents" to string type
	fmt.Println(reflect.TypeOf(contents))
	fmt.Println(contents)
	fmt.Println(string(contents))

	f, err := os.Open("file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// func NewScanner(r io.Reader) *Scanner
	s := bufio.NewScanner(f)
	// func (s *Scanner) Scan() bool
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
