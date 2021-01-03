package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waiter sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("CPU Count:", runtime.NumCPU())
	incrementor := func() {
		defer waiter.Done()
		mutex.Lock()
		temp := counter
		fmt.Println("Incrementing...", temp)
		temp++
		counter = temp
		mutex.Unlock()
	}
	iter := 30
	waiter.Add(30)
	for i := 0; i < iter; i++ {
		go incrementor()
	}
	waiter.Wait()
	fmt.Println(counter)
	fmt.Println("Number of routines in Use:", runtime.NumGoroutine())
}
