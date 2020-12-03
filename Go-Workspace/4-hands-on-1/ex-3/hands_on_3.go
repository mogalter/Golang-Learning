package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprint(y, " is ", fmt.Sprint(42)+".", " It's definitely ", z)
	fmt.Println(s)
}
