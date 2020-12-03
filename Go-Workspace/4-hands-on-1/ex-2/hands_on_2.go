package main

import "fmt"

func main() {
	// assigned values are 0, '', and False (defaults for their respective type)
	var x int
	var y string
	var z bool
	// single line prints
	// single print statement
	fmt.Printf("x has a type of %T and a value of <%v>\n", x, x)
	fmt.Printf("y has a type of %T and a value of <%v>\n", y, y)
	fmt.Printf("z has a type of %T and a value of <%v>\n", z, z)
}
