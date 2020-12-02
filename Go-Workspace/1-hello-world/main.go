package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	foo()
	n, e := fmt.Println("End of", "main", "package")
	fmt.Println(n, e)
}

func foo() {
	fmt.Println("Foo", "Foo", "Foe", "Fum")
}
