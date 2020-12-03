package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	// single line prints
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// single print statement
	fmt.Println(y, "is", fmt.Sprint(42)+".", "It's definitely", z)
}
