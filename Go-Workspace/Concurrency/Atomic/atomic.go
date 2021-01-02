package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// atomic allows us to avoid race conditions
// template code copied from race condition
var counter int64

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
			// mutexes prevent a race condition issue
			// by locking access to code while routine is executing
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
		}()
		fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
	}
	waiter.Wait()
	fmt.Println(counter) // should be 50 were it not for race conditions
}
