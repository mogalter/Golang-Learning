package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	x := []byte("qwerty")
	bs, err := bcrypt.GenerateFromPassword(x, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}
	fmt.Println("Correct hash:", bcrypt.CompareHashAndPassword(bs, x))
}
