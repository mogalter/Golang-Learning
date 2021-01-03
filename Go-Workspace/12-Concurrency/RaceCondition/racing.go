package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

func main() {
	fmt.Println("Number of CPUs in Use:", runtime.NumCPU())
	fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	gs := 50
	waiter := sync.WaitGroup{}
	waiter.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer waiter.Done()
			// this will create a race condition
			// we retrieve the counter val and yield
			// this lets other routines grab the old counter value.
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
		}()
		fmt.Println("On iteration", i)
		fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	}
	waiter.Wait()
	fmt.Println(counter) // should be 50 were it not for race conditions
}
