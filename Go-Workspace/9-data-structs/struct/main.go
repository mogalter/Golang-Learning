package main

import "fmt"

// define a struct
type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("Hello!")
	// sample instatiation of struct
	kevin := person{
		first: "Kevin",
		last:  "Huang",
		age:   24,
	}

	fmt.Println(kevin.first, kevin.last)
}
