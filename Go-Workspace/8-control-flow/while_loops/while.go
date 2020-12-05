package while

import "fmt"

// LoopWhile returns a greeting for the named person.
func LoopWhile(repeat int) int {
	// Return a greeting that embeds the name in a message.
	fmt.Println("Executing while loop!")
	i := 0
	for i < repeat {
		fmt.Println(i)
		i++
	}
	return i
}
