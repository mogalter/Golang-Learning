package main

// fan in takes values from many channels and places into 1 channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	worker := make(chan int)
	aggregator := make(chan int)
	go populate(worker)
	go fanIn(worker, aggregator)
	for v := range aggregator {
		fmt.Println(v)
	}
	fmt.Println("Done")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanIn(worker, aggregator chan int) {
	var waiter sync.WaitGroup
	//throttling, fixed routine count to 10
	const routines = 10
	waiter.Add(routines)
	// prepare to trip a go routine!
	// go func(randVar int) {
	// 	aggregator <- heavyLifting(randVar)
	// 	// indicate this go routine is done
	// 	waiter.Done()
	// }(v)
	//}
	for i := 0; i < routines; i++ {
		// for each count in routines, launch following routine:
		// here, we will have 10 concurrent routines running!
		go func() {
			for v := range worker {
				func(v2 int) {
					aggregator <- heavyLifting(v2)
				}(v)
			}
			waiter.Done()
		}()
	}
	// wait until we place all our values onto aggregator channel
	waiter.Wait()
	close(aggregator)
}

func heavyLifting(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(200)
}
