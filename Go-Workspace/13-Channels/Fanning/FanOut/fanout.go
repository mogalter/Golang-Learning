package main

// fan in takes values from many channels and places into 1 channel

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)
	go send(even, odd)
	go receive(even, odd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func send(e, o chan<- int) {
	// sender chans
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// q <- true
	close(e)
	close(o)
}

func receive(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
