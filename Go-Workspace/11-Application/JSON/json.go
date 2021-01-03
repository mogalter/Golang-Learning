package main

import (
	"encoding/json"
	"fmt"
)

// Thief - struct to represent a thief
type Thief struct {
	Name   string   `json:"Name"`
	Weapon string   `json:"Weapon,omitempty"`
	Allies []string `json:"Allies"`
}

func main() {
	phantom := Thief{
		Name:   "Phantom",
		Allies: []string{"Mercedes", "Evans"},
	}
	//json.marshal turns an slice -> json-like object (slice of bytes)
	//json.unmarshal unpacks a json object (a slice of bytes). Takes in an address to store info in
	var thieves Thief
	temp, _ := json.Marshal(phantom)
	fmt.Printf("%v - %T\n", string(temp), temp)
	err := json.Unmarshal(temp, &thieves)
	fmt.Println(thieves, err)
}
