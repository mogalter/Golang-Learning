package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// creators of go feel try-catch-finally results in convoluted code
// defer, panic, recover will be used for exceptional condition
// no assertions in Go either!
// another way to think of it: another type!

type person struct {
	name string
}

func (p person) Error() string {
	return p.name + "has stopped functioning!"
}

func sampleWriter() {
	// example
	f, err := os.Create("aloha.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Ohallo")
	io.Copy(f, r)
}

func sampleReader() {
	f, err := os.Open("aloha.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	fmt.Println(string(bs))
}

func main() {
	fmt.Println("Error Handling")
	// error creation
	err := errors.New("Woopsies")
	fmt.Println(err.Error())
	sampleWriter()
	sampleReader()
}
