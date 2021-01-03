package main

import (
	"fmt"
	"runtime"
)

func passTwo() {
	// usage of buffered channels; below init allows 1 value to sit on channel
	// we're blocked until the channel capacity is emptied!
	c := make(chan int, 1)
	c <- 32
	fmt.Println(cap(c), <-c)
}

func main() {
	fmt.Println("Executing on:", runtime.GOARCH+"-"+runtime.GOOS)
	// it's important to note that channels block
	// we can't do anything with a channel until we can send and recieve infomation at once
	c := make(chan int) // a unbuffered channel that can take in int values
	go func() {
		c <- 42
	}() // places 42 on the channel
	// extracts number 42 and outputs it from channel - the next line is blocked till func is done.
	fmt.Println(<-c)
	passTwo()
}
