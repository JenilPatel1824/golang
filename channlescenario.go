package main

import "fmt"

func main() {
	ch := make(chan int) // Create a channel

	// This will block because there is no sender
	//val := <-ch
	//close(ch)
	fmt.Print("trying to recive: ")
	fmt.Println(<-ch)
	val, open := <-ch
	fmt.Println(val, open)
	//ch <- fmt.Println("Received:")
	// main function ends without sending, receiver will be blocked forever
}
