package main

import "fmt"

type superSaiyan int

var x superSaiyan

func main() {
	fmt.Println("Value of Super Saiyan", x)
	fmt.Printf("Value of Super Saiyan 'x' is %T\n", x)
	x = 42
	fmt.Println("Value of Super Saiyan", x)
}
