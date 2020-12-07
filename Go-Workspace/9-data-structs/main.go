package main

import "fmt"

func slicing() {
	// slice creation
	// composite literal: a way to create different values of types
	// x := type{values}
	// Slices lets us group together values of the same time
	x := []int{4, 5, 6, 42}
	// append func takes in a slice of type T and values we want to append
	x = append(x, 2, 3, 1)
	// fmt.Printf("%T, %v\n", x, x)
	// slice range looping - Python enumerate() equivalent
	for i, v := range x {
		fmt.Println(i, v)
	}
}

func slicingASlice() {
	x := []int{4, 5, 6, 42}
	// similar to python: x[start:end]
	// end is excluded from range, if no end, we simply go to en
	// this is how we delete from a slice as well
	x = x[1:]
	fmt.Println(x)
}

func appendSlice() {
	x := []int{4, 5, 6, 42}
	y := []int{77, 8, 9, 100}
	// regular append
	x = append(x, 11, 120, 17)
	// y... unfurls "y"
	// ... tells the compiler to 'unpack' y into a list of args and append it
	x = append(x, y...)
	fmt.Println(x[:])
}

func main() {
	// Array creation
	// var <var name> [size of array] <type in array>
	// "Arrays are used to build slices; Go discourages using Arrays and recommends slices"
	var x [5]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	fmt.Printf("%T, %v\n", x, x) // will print [5]int <contents of x>
	//slicing()
	//slicingASlice()
	appendSlice()
}
