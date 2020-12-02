package main

import "fmt"

// we cannot have a non-declaration statement outside function body
// x := 44 <-- cannot do this outside of function body
// In short var = global, := is local

var y = 45

// Here, we declare a variable z and specify it's type (int)
// If no value is provided, we use a default value
var z int

// Tip: use short declaration as much as possible
func main() {
	// variable declaration + value assignment, short declaration
	// ** of a certain type
	x := 42
	fmt.Println(x)
	fmt.Println(y)
}
