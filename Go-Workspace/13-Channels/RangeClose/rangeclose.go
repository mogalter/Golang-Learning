package main

import "fmt"

func foo(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- i
	}
	fmt.Println("Sender channel closed")
}

// receive, will block until channel receives value
func bar(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan int)
	// sender
	go foo(c)
	// for range loops over channel until it's closed
	// if no close: we go into an deadlock since we're still waiting for c
	for num := range c {
		fmt.Println("Pulled value from channel:", num)
	}
	fmt.Println("Ending")
}
