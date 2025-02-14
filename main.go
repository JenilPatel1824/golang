package main

import (
	"fmt"
	"log"

	"golang/xyz"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("xyz: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Set logging format

	// Request a greeting message.
	message, err := xyz.Hello("")

	fmt.Printf("%T", err) // If an error was returned, print it to the console and
	// exit the program.
	//if err != nil {
	//	log.Fatal(err)
	//}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	var x int

	fmt.Println(x)
}
