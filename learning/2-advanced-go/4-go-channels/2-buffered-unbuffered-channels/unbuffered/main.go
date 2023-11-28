package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go send(ch)
	fmt.Println("\nblocking send goroutine...")
	fmt.Printf("channel length: %v\n", len(ch))
	go receive(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan string) {
	ch <- "message"
}

func receive(ch chan string) {
	time.Sleep(time.Second * 1)
	fmt.Println("send goroutine unblocked")
	fmt.Printf("channel length: %v\n", len(ch))
	fmt.Printf("received: %v\n", <-ch)
}
