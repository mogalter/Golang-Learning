package main

import (
	"fmt"
)

// send
func foo(c chan<- int) {
	c <- 42
}

// receive, will block until channel receives value
func bar(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)
	// sender
	go foo(c)
	// receiver - we do not put a go at bar
	// why? because we want bar(c) to hold until foo(c) is done
	bar(c)
	fmt.Println("Ending")
}
