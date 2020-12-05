package main

import "fmt"

func main() {
	// untyped constants have their types assigned when referenced
	fmt.Println("Untyped-Typed Point of Entry")
	// typed
	const (
		a int = 1.0
	)
	// Untyped
	const (
		b = 2.0
	)
	fmt.Printf("Typed: %T %v, Untyped: %T %v", a, a, b, b)
}
