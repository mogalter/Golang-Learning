package main

import "fmt"

func main() {
	// hex = %x, %b = bin, %d = dec
	a := 20
	fmt.Printf("Original: %v, Bin: %b, Hex %x, Dec %d\n", a, a, a, a)
	a = a << 1
	fmt.Printf("Shifted: %v, Bin: %b, Hex %x, Dec %d\n", a, a, a, a)
}
