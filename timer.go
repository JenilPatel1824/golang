package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//var wg sync.WaitGroup

	var wg sync.WaitGroup
	fmt.Println("starting timer and perodic ticker")

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		timer := time.NewTimer(3 * time.Second)

		<-timer.C

		fmt.Println("timer of goroutin1")
		wg.Done()
	}(&wg)

	//wg.Add(1)
	fmt.Println("added 2 to wg")

	go func() {
		fmt.Println("starting periodic task")
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop() // Make sure to stop the ticker when done

		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticked!")
			}
		}
		fmt.Print("free")
	}()
	fmt.Println("waiting for go routine to done")

	fmt.Println("meanwhile main is  free")
	wg.Wait()

}
