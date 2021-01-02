package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waiter = sync.WaitGroup{}

func foo() {
	fmt.Println("FOOOOOO")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	waiter.Done() // decrement wait counter by one
}

func bar() {
	fmt.Println("BARRRRR")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func init() {
	// runs before main (once)
}

func main() {
	// put something in a go routine simply by putting go in front of it.
	// all go routines execute once waiter's counter hits 0
	waiter.Add(1)
	go foo()
	bar()
	fmt.Println("Number of CPUs in Use:", runtime.NumCPU())
	fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	// holds up the code here until wait hits 0
	waiter.Wait()
	// concurrency vs parallelism?
	// concurrency - a design pattern where code can potentially run in parallel
	// does not guarantee it runs in parallelism
	// we can run things in parallel if we have more than 1 CPU
}
