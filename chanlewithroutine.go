package main

import (
	"fmt"
	"time"
)

func timeticker() {
	fmt.Println("time ticker called it will just print time at each second")
	//ticker := time.NewTicker(1 * time.Second)
	//
}
func sender1(ch chan string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		ch <- "Sender 1: Data sent every 5 seconds"
	}
}

func sender2(ch chan string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		ch <- "Sender 2: Data sent every 5 seconds"
	}
}

func receiver(ch1, ch2 chan string) {
	for {
		select {
		case msg := <-ch1:
			fmt.Println("Received from sender1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from sender2:", msg)
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sender1(ch1)
	go sender2(ch2)
	go receiver(ch1, ch2)

	// Keep the main function alive
	select {}
}
