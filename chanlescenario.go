package main

import (
	"fmt"
	"time"
)

func main() {

	var ch = make(chan int)
	//close(ch)
	//fmt.Print(<-ch)
	//fmt.Print(<-ch)

	go func() {
		ch <- 30

		defer func() {
			fmt.Println("njndeke")

			if r := recover(); r != nil {
				fmt.Println("rrcvered    ", r)
			}
			time.Sleep(1 * time.Millisecond)
		}()

		defer panic("aaa    ")
		defer fmt.Print("raj      ")

		fmt.Println("done send")

	}()
	fmt.Print(<-ch)

	//defer ticker.Stop()
	//
	//go func() {

	//}()
	//time.AfterFunc(5*time.Second, func() {
	//	fmt.Print("done")
	//})
	//fmt.Println("main free")
	//
	//ticker := time.NewTicker(2 * time.Second)
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	ticker.Stop()
	//}()
	//
	//for t := range ticker.C {
	//	fmt.Println("Heartbeat at:", t)
	//}
	//
	//fmt.Println("reciving")
	//x := <-ch
	//fmt.Println(x)

	//ch <- 1

	//time.Sleep(10 * time.Second)
	//time.AfterFunc(5000*time.Second, func() {
	//	fmt.Println("Retrying API call now!")
	//})
	//fmt.Println("timmer started")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recovered: ", r)
	//	}
	//}()
	//
	////go func() {
	////	time.Sleep(1 * time.Hour)
	////}()
	//
	//var wg = sync.WaitGroup{}
	//wg.Add(1)
	//ch := make(chan int)
	//go func(wg *sync.WaitGroup) {
	//	//time.Sleep(10 * time.Millisecond)
	//	ch <- 2
	//
	//	ch <- 3
	//
	//	//fmt.Println("sent 2 seping for 5 sec ")
	//	//time.Sleep(5 * time.Second)
	//	fmt.Println("after sleep")
	//	//ch <- 5
	//
	//	//close(ch)
	//
	//	//close(ch)
	//}(&wg)
	//
	//go func() {
	//	fmt.Println(<-ch) // Blocks forever
	//	fmt.Println(<-ch) // Blocks forever
	//	// Blocks forever
	//	//
	//	//fmt.Println(<-ch) // Blocks forever
	//
	//	// Blocks forever
	//	//
	//	//fmt.Println(<-ch) // Blocks forever
	//
	//}()
	//wg.Wait()
	//
	////ch2 := make(chan int, 1)
	////ch2 <- 1
	////ch2 <- 2
	////ch2 <- 3
	////close(ch2) // Close channel
	////
	////for val := range ch2 {
	////	fmt.Println(val) // Output: 1 2 3
	////}
	//
	////fmt.Println("trying to recive: 2")
	////fmt.Println(<-ch)
	////fmt.Println("trying to recive: 3")
	////fmt.Println(<-ch)
	////fmt.Println("trying to recive: 4")
	////fmt.Println(<-ch)
	////fmt.Println("trying to recive: 5")
	////fmt.Println(<-ch)

	//wg.Wait()
	fmt.Println("after wait over")

}
