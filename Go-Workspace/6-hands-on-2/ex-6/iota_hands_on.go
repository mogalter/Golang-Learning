package main

import "fmt"

func main() {
	// untyped constants have their types assigned when referenced
	fmt.Println("Iota Point of Entry")
	// typed
	const (
		a = iota + 2020
		b
		c
		d
	)
	fmt.Println("Years:", a, b, c, d)
	// Untyped
}
