package main

import "fmt"

// define a struct
// fields in person are promoted in secretAgent!
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person // embed person into secret agent
	ltk    bool
}

func main() {
	fmt.Println("Hello!")
	// sample instatiation of struct
	kevin := person{
		first: "Kevin",
		last:  "Huang",
		age:   24,
	}
	fmt.Println(kevin)
	finalAgent := secretAgent{
		person: kevin,
		ltk:    true,
	}
	finalAgent.age = 32
	// we can call fields of embedded structs directly
	// in the case there are name collisions, we can specify by mentioning the class directly
	// e.g. if secretAgent also had a name field, we can get person's by: secretAgent.person.name
	fmt.Println(finalAgent) // final agent's instance of "Kevin" is a deep copy!
	fmt.Println(kevin)
}
