package main

import "fmt"

// what is select? - basically a switch statement
// pulls value off channel whenever available. blocks

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(even, odd, quit)
	receive(even, odd, quit)
  close(even)
  close(odd)
}

func send(e, o, q chan<- int) {
	// sender chans
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	fmt.Println("Hello")
  close(q)
	// q <- true
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even:", v)
		case v := <-o:
			fmt.Println("From odd:", v)
		case i, ok := <-q:
      // will execute last bc this will be empty till loop exit!
      if !ok {
        // this will execute if channel is closed.
        // ok = false denotes a closed channel
        // an empty channel would yield: 0 + false
        fmt.Println("From quit:", i, ok)
      } else {
        fmt.Println("Pulled", i)
      }
			return
		}
	}
}
