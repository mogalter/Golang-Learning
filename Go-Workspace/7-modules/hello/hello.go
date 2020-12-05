package main

// Sample package for use in importing
import (
	"fmt"
	// we will import example.com/greetings
	"example.com/greetings"
)

// replace example.com/greetings => ../greetings
// this line of code in go.mod will tell the compiler to replace
// example.com/greetings with another directory

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
