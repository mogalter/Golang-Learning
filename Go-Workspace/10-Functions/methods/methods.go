package main

import "fmt"

// Node - the underlying type for a linked list
type Node struct {
	Data int
	Next *Node
}

// Insert method for a linked list - note that the receiver ties this to the node struct!
func (root *Node) Insert(val int) {
	for root.Next != nil {
		root = root.Next
	}
	root.Next = &Node{val, nil}
}

// PrintList method prints the entire list
func (root *Node) PrintList() {
	for root != nil {
		fmt.Print(root.Data, "->")
		root = root.Next
	}
	fmt.Println("END")
}

func main() {
	sample := Node{3, nil}
	sample.Insert(10)
	sample.PrintList()
}
