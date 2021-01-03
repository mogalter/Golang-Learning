package main

import (
	"fmt"
	"runtime"
)

func senderOnly() {
	c := make(chan<- int, 2)
	c <- 30
	c <- 31
	fmt.Printf("%v, %T\n", c, c)
}

func main() {
	// we can make channels we can only receive off of/write onto
	fmt.Println("Executing on:", runtime.GOARCH+"-"+runtime.GOOS)
	// send only channel: make(chan <- int), cannot pull values off channel
	// recieve only: make(<-chan int), cannot place values on channel
	// key to remembering: read left to right
	c := make(chan int, 2) // make a channel of ints with cap 2
	cr := make(<-chan int, 2)
	cs := make(chan<- int, 2)
	// cannot assign directional to bidirectional
	// c = cr/cs <-- will not compile
	// however, we can go from bidirectional to directional
	fmt.Println(cr == c)
	cr = c
	fmt.Println(cr == c)
	fmt.Printf("c\t%T\ncr\t%T\ncs\t%T\n", c, cr, cs)
}
