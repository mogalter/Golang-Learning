package main

import "fmt"

func variadicType(x ...int) {
	// this function utilizes a variadic parameter!
	// what is a variadic parameter?
	// its a parameter that accepts infinite vals and those get stored in a slice
	// this type of parameter must be the FINAL parameter!
	fmt.Printf("In Var Type: %T, %v\n", x, x)
}

func deferTheDefer(switches int) bool {
	defer variadicType(1, 2, 3, 4, 5)
	if switches == 1 {
		fmt.Println("Ret true")
		return true
	}
	// defers the execution of a function
	// till current function returns or ends!
	//  e.g. defer file close when file reading ends.
	fmt.Println("Ret false")
	return false
}

func main() {
	// func (r receiver) identifier(parameter(s)) (return(s)) { code }
	// receiver refers to what should receive the function (who the func belongs to)
	sample := []int{1, 2, 3, 4, 5}
	// we unfurl sample and pass it in... as a stream of infinite values
	//... before a type refers to infinite vars of that type
	//... after a slice refers to unfurling the slice
	variadicType(sample...)
	deferTheDefer(1)
	deferTheDefer(7)
}
