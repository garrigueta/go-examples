package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func(ch chan string) {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			msg := "message" + strconv.Itoa(i)
			ch <- msg
			fmt.Printf("SEND goroutine: %v\n", msg)
		}
	}(ch)

	go func(ch chan string) {
		for val := range ch {
			fmt.Printf("RECV goroutine: %v\n", val)
		}
	}(ch)

	fmt.Println("MAIN goroutine: sleeping!")
	time.Sleep(time.Millisecond * 100)
	val, ok := <-ch
	fmt.Printf("values returned from the read operation 'val, ok := <-ch' are: %q, %v\n", val, ok)
	fmt.Println("MAIN goroutine: done!")
}
