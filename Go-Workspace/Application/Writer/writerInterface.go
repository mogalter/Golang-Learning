package main

import "fmt"

// Decode iterates over JSON
// Encode writes to JSON via a stream
// Writer interface is an interface for types that implement Write([]bytes)
// e.g. OS file object is also interface Writer because it implements Write()

func main() {
	fmt.Println("Writer Interface")
}
