package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// context helps prevent goroutine leaks
// context is a package that makes it easy to pass
// request-scoped vals (e.g. end user auth, cookies), cancellations, deadlines to all goroutines
// handling a request!

// func main() {
//   ctx := context.Background()
//
//   fmt.Println("context:\t", ctx)
//   fmt.Println("context err:\t", ctx.Err())
//   fmt.Printf("context type:\t%T\n", ctx)
//   fmt.Println("----------")
//
//   ctx, cancel := context.WithCancel(ctx)
//
//   fmt.Println("context:\t", ctx)
//   fmt.Println("context err:\t", ctx.Err())
//   fmt.Printf("context type:\t%T\n", ctx)
//   fmt.Println("cancel:\t\t", cancel)
//   fmt.Printf("cancel type:\t%T\n", cancel)
//   fmt.Println("----------")
//
//   // cancel tells op to abandon it's work
//   // does not wait for work to stop
//   cancel()
//
//   // once cancel is called, err will be marked as: context canceled
//   fmt.Println("context:\t", ctx)
//   fmt.Println("context err:\t", ctx.Err())
//   fmt.Printf("context type:\t%T\n", ctx)
//   fmt.Println("cancel:\t\t", cancel)
//   fmt.Printf("cancel type:\t%T\n", cancel)
//   fmt.Println("----------")
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			// done is triggered when cancel is called
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutines 3:", runtime.NumGoroutine())
}
