package main

import "fmt"

func main() {
	fmt.Println("Entrypoint")
	a := 2 == 2
	b := 1 <= 3
	fmt.Println(a, b)
	c := 1 >= 3
	d := 1 != 1
	e := 1 < 2
	f := 1 > 2
	fmt.Println(c, d, e, f)
}
