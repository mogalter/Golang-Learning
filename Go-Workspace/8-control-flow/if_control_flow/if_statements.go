package ifStatements

import "fmt"

// CheckNum does a check on numerical input and gives an output accordingly
func CheckNum(x int) {

	if x == 1 {
		fmt.Println("Wow, this is a one!")
	} else if x == 2 {
		fmt.Println("Only a two?")
	} else {
		fmt.Println("Hmmm, we got a", x)
	}
	// in the case we only want to use x once, we can do:
	// x := <init val>; <comparison>
	// sample check

	if x := 9001; x > 9000 {
		fmt.Println("It's over 9000!!")
	} else if x == 42 {
		fmt.Println("It's a mere 42...")
	} else {
		fmt.Println("Power Level", x)
	}

}
