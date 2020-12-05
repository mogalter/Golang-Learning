package main

import "fmt"

// define iotas
const (
	a int = iota * 10 // <var name> <type> for declaration,
	// if we do multiplication, that will be the increment!
	b
	c
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c)
}
