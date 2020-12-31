package main

import "fmt"

// a callback is when we pass in a function as an argument
// "Functional Programming"

func sum(x ...int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}

func even(f func(x ...int) int, nextArr ...int) int {
	evens := []int{}
	for _, num := range nextArr {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return f(evens...)
}

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 20}
	s := sum(i...)
	fmt.Printf("%T, %v\n", s, s)
	fmt.Println("Even sums", even(sum, i...))
}
