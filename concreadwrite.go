package main

import (
	"fmt"
	"time"
)

func main() {
	myMap := make(map[int]string)

	// Goroutine 1: Writing to the map
	go func() {
		for i := 0; i < 1000; i++ {
			myMap[i] = fmt.Sprintf("Value %d", i)
		}
	}()

	// Goroutine 2: Reading from the map
	go func() {
		//time.Sleep(5 * time.Microsecond)
		for i := 0; i < 50; i++ {
			fmt.Println(myMap[i]) // Concurrent read while writing
		}
	}()

	time.Sleep(time.Second) // Give goroutines time to execute
	fmt.Println(myMap)
}
