package main

import "fmt"

// LRUCache - for use with memoization
func LRUCache(comp func(int) int) func(int) int {
	// the state is kept for seen
	seen := map[int]int{}
	return func(num int) int {
		if _, ok := seen[num]; !ok {
			seen[num] = comp(num)
		}
		return seen[num]
	}
}

// this keeps state, calling this over and over will increment i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Fibonacci
func fibonacci(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	newFib := LRUCache(fibonacci)
	fmt.Println(newFib(20))
}
