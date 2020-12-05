package custom_types

import "fmt"

// make type hotdog with an underlying type in int
type hotdog int

var a int

// create variable 'b' with name of hotdog
var b hotdog

func Custom() {
	fmt.Println("Entrypoint of the script")
	fmt.Printf("Variable b, has a type %T\n", b)
	// since Go is status, we can't set a=b. a is an int but b is a hotdog
}
