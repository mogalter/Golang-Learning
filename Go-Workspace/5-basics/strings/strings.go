package main

import "fmt"

func main() {
	s := "hello world"
	bs := []byte(s)
	fmt.Println("s")
	for i := 0; i < len(bs); i++ {
		fmt.Printf("Index %v, UTF-8 Value %U\n", i, bs[i])
	}
}
