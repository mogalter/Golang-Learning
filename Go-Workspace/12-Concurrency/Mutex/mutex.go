package main

import (
	"fmt"
	"runtime"
	"sync"
)

// from race condition doc

var counter int

func main() {
	fmt.Println("Number of CPUs in Use:", runtime.NumCPU())
	fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	gs := 50
	mutex := sync.Mutex{}
	waiter := sync.WaitGroup{}
	waiter.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer waiter.Done()
			// mutexes prevent a race condition issue
			// by locking access to code while routine is executing
			mutex.Lock()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			mutex.Unlock()
		}()
		fmt.Println("On iteration", i)
		fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	}
	waiter.Wait()
	fmt.Println(counter) // should be 50 were it not for race conditions
}
