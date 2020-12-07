package main

import (
	"fmt"

	"example.com/ifStatements"
	"example.com/initCondPost"
	"example.com/switching"
	"example.com/while"
)

func main() {
	x := 7
	fmt.Println("Repeated:", x)
	while.LoopWhile(x)
	initCondPost.LoopFor()
	ifStatements.CheckNum(x)
	switching.SwitchOn(7)
	switching.SwitchOn(100)
}
