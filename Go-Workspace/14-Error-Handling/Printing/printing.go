package main

import (
	"fmt"
	"log"
	"os"
)

// Fatal = exit!
// Panic = log and recover (deferred programs run)

// PanicSample provides an example of panic()
func PanicSample() {
	bar, err := os.Open("bar.txt")
	defer func() {
		bar.Close()
		fmt.Println("File closed.")
	}()
	if err != nil {
		// log.Println("Oh boy, no file!") -> prints <time> <passed in str>
		// panic(err)
		log.Panicln("Oh boy, no file found... shutting down!")
	}
}

func main() {
	fmt.Println("Hello")
	//example
	// set a file output location
	logOutput, err := os.Create("log.txt")
	if err != nil {
		log.Panicln("Unable to create log file!")
	}
	defer logOutput.Close()
	log.SetOutput(logOutput)
	// write error to log :)
	PanicSample()
}
