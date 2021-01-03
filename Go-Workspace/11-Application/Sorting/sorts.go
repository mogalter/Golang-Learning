package main

import (
	"fmt"
	"sort"
)

// Person - describes a person
type Person struct {
	Name string
	Age  int
}

// ByAge is an array of Person for use in sorting
// This type will match the sort.Interface interface implementation
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	fmt.Println("Sorting")
	// int slice is a defined type in sort package (basically a int[])
	x := sort.IntSlice{1, 7, -10, 2, 20, -100}
	sort.Ints(x)
	fmt.Println(x, sort.IsSorted(x))
	// custom sorting below
	p1 := Person{"James", 32}
	p2 := Person{"Alex", 16}
	p3 := Person{"Kevin", 24}
	y := []Person{p1, p2, p3}
	sort.Sort(ByAge(y))
	fmt.Println(y)
}
