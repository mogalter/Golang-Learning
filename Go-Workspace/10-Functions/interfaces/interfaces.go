package main

import (
	"fmt"
)

// interfaces let us do polymorphism!
// interfaces allow us to define behavior.
// example below - idea: a value can be of more than 1 type

type agent struct {
	name        string
	age         int
	killLicense bool
}

type person struct {
	name string
}

func (secretAgent agent) speak() {
	fmt.Println("I am", secretAgent.name, "and I have come for your life!")
}

func (ordinary person) speak() {
	fmt.Println("I am", ordinary.name)
}

func discuss(h human) {
	fmt.Print("Human, discuss(): ")
	// switching on type + assertion
	// assertion is saying 'I KNOW this is type "X"'
	switch h.(type) {
	case person:
		fmt.Println("Discussing", h.(person).name)
	case agent:
		fmt.Println("Does your agent have a LTK?", h.(agent).killLicense)
	}
}

// in this case, whatever object can speak is a human
// e.g secretAgent is both human and agent at once
type human interface {
	speak()
}

func main() {
	fmt.Println("hello")
	agentOne := agent{
		"Ben", 20, true,
	}
	personOne := person{"Dr. Yes"}
	fmt.Println(personOne, agentOne)
	agentOne.speak()
	personOne.speak()
	discuss(agentOne)
	discuss(personOne)
	// assetion - asserting something is a certain type.
}
