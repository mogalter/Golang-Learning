package main

import "fmt"

// closures help maintain a separate scope within a function
// enclose variables in a functional scope
// "A closure is a mechanism that allows a function
//  to remember a scope in which it was defined"
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	// incrementor holds it's own scope
	ticker := incrementor()
	fmt.Println(ticker())
	fmt.Println(ticker())
	fmt.Println(ticker())
	fmt.Println(ticker())
	fmt.Println(ticker())
}
