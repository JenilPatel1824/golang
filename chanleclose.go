// Go program to illustrate how
// to close a channel using for
// range loop and close function
package main

import (
	"fmt"
	"time"
)

// Function
func myfun(mychnl chan string) {

	close(mychnl)

	for v := 0; v < 4; v++ {
		time.Sleep(3 * time.Second)
		mychnl <- "GeeksforGeeks"
	}
	time.Sleep(3 * time.Second)
	close(mychnl)
}

// Main function
func main() {

	c := make(chan string)

	mesg := <-c

	go myfun(c)
	fmt.Println("reciving")
	fmt.Println("Recived", mesg)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
