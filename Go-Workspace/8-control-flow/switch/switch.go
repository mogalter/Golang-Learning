package switching

import "fmt"

// SwitchOn will take in a value and return the case output
func SwitchOn(switchVal int) int {
	fmt.Println("Switching", switchVal)
	// Explicitly writing out the conditionals
	// Missing switch expressions are turned into True!
	// switch {
	// case switchVal == 10:
	// 	fmt.Println("Indeed, this is a 10/10")
	// 	// if we remove the return and fallthrough, we'll hit switch == 7
	// 	return 1
	// case switchVal == 7:
	// 	fmt.Println("Lucky 7s!!")
	// 	return 777
	// default:
	// 	fmt.Println("Not sure what you just gave me...")
	// 	return -1
	// }
	// if you fallthrough, everything after a fallthrough gets dumped out
	// usage of fallthrough is discouraged

	// this usage of switch omits the need to write out conditionals
	// whatever is placed after case will be checked by switchVal
	switch switchVal {
	case 10, 100, 1000: // we can also string together cases if they share a same result
		fmt.Println("Indeed, this is a 10/10")
		// if we remove the return and fallthrough, we'll hit switch == 7
		return 1
	case 7:
		fmt.Println("Lucky 7s!!")
		return 777
	default:
		fmt.Println("Not sure what you just gave me...")
		return -1
	}
}
