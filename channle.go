package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2) // Unbuffered channel

	go func() {
		fmt.Println("Sending 1 (Blocks until received)")
		ch <- "Message 1" // BLOCKS until main receives it
		fmt.Println("Sent 1")

		fmt.Println("Sending 2 (Blocks until received)")
		ch <- "Message 2" // BLOCKS until main receives it
		fmt.Println("Sent 2")
	}()

	time.Sleep(5 * time.Second) // Simulating delay before receiving

	// Receiving messages
	fmt.Println("Receiving 1:", <-ch) // Unblocks first send
	fmt.Println("Receiving 2:", <-ch) // Unblocks second send
}
