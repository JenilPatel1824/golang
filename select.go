package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(5 * time.Second): // Timeout case
		fmt.Println("No response, timing out")

	default:
		fmt.Println("def")
	}
	fmt.Println("free")
}
