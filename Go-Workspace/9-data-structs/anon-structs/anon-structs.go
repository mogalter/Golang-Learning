package main

import "fmt"

// anonymous field: field with no name
// format for entry in structs identifier list type or anon field (tag)
// in the case of an anon field, it should be a type
// embedded field must be a type or a pointer to a type name. e.g. *Person
// type name acts as the field name

var x, y int

// embedded type in a struct.
type numCruncher struct {
	int
	y int
}

func (n *numCruncher) crunch() int {
	n.int = n.int + n.y
	return n.int
}

func main() {
	x, y = 20, 11
	// we can do an assignment in a pythonic way!
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
	fmt.Println("Anon structs.")
	n := numCruncher{
		int: x,
		y:   y,
	}
	fmt.Println(n)
	n.crunch()
	fmt.Println(n)
	// anonymous structs
	// lets you define a struct on the go and use it :)
	person := struct {
		name string
		age  int
	}{
		name: "James Bond",
		age:  32,
	}
	fmt.Println(person)
}
