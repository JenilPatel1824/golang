package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Waiting 5 seconds before retrying API call...")

	//ticker := time.NewTicker(2 * time.Second)
	//defer ticker.Stop()
	//
	//go func() {
	//	for t := range ticker.C {
	//		fmt.Println("Heartbeat at:", t)
	//	}
	//}()

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Retrying API call now!")
	})

	time.Sleep(6 * time.Second) // Wait to see the output

	time.Sleep(7 * time.Second)
}
