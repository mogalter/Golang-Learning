package main

import "fmt"
import "../3-types/zeros"
import "../3-types/custom-types"

// Go is statically typed, a typed variable must remain that type
// Here, we explicitly configure var y with a type of "int": var y int = 42
var y = 42

// Raw string literals are constants made by concactenanting a seq. of chars
// Denoted by ticks: ``
var z = `Raw String Literal`

func main() {
	// printf lets us swap in formatted vars into a string
	x := 2
	fmt.Printf("Ohallo, we have y with a type of %T and value of %d\n", y, y)
	fmt.Println(x)
	fmt.Println(z)
	// Exported funcs must have capitalized function names for access
	zeros.Hello()
	// Sprint will print to a string for assignment
	// %v will print the normal value
	s := fmt.Sprintf("Ohallo, we have y with a type of %T and value of %d\n", y, y)
	fmt.Print("Output of sprintf: ")
	fmt.Println(s)
	custom_types.Custom()
}
