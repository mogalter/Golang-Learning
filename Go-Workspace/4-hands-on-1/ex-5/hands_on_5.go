package main

import "fmt"

type superSaiyan int

var x superSaiyan
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("x's Type: %T\n", x)
	x = 24
	fmt.Println("x now has a value of", x)
	y = int(x)
	fmt.Printf("y's Type: %T\n", y)
	fmt.Println("y now has a value of", y)
}
