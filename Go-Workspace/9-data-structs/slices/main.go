package main

import "fmt"

func slicing() {
	// slice creation
	// composite literal: a way to create values for structs
	// LiteralType LiteralValue: e.g. [] int
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

func sliceDelete() {
	x := []int{4, 5, 6, 42, 100, 55}
	// this will append everything after 4th index to post-2nd index
	// recall append(<structure to append to>, <to append>...)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func makeSlicing() {
	// make(composite literal, cur length, max capacity)
	// the following makes an array w. a len of 10 and contains 10 zeros
	x := make([]int, 10, 100)
	fmt.Println(x, len(x), cap(x))
	// manually setting values: x[index] = <value>
	// we can also still append
	x = append(x, []int{3, 10, 9}...)
	fmt.Println(x, len(x), cap(x))
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
	sliceDelete()
	makeSlicing()
}
