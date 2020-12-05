package main

import (
	"fmt"

	"example.com/ifStatements"
	"example.com/initCondPost"
	"example.com/while"
)

func main() {
	x := 7
	fmt.Println("Repeated:", x)
	while.LoopWhile(x)
	initCondPost.LoopFor()
	ifStatements.CheckNum(x)
}
