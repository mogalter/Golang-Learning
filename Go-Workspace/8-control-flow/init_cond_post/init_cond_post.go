package initCondPost

import "fmt"

// LoopFor displays a sample loop
func LoopFor() {
	fmt.Println("In the Init, Condition, Go! Package")
	// init var; condition to continue loop; what to do after loop action
	// these are the loop basics for Golang
	for i := 10; i < 100; i += 20 {
		fmt.Println("For Loop Result:", i)
	}
}
