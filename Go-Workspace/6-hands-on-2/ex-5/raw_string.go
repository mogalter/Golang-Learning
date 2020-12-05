package main

import "fmt"

func main() {
	// Special characters are not taken into account in raw string
	a := `Hello World\n`
	fmt.Println(a)
}
