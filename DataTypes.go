package main

import (
	_ "expvar" // Exposes metrics
	"fmt"
	"log"
	"net/http"
	"time"
)

var myvariable1 int
var myvariable2 string
var myvariable3 float64

// Display the zero-value of the variables

func main() {

	go func() {
		log.Println(http.ListenAndServe(":8081", nil)) // Metrics available at :8081
	}()
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvariable2)

	fmt.Printf("The value of myvariable3 is :",
		myvariable3)

	str := "hello"
	fmt.Println("f")
	fmt.Println(str[0])

	if myvariable1 == 0 {
		fmt.Println("m==0")
	} else if myvariable1 >= 0 {
		fmt.Println("else if")
	}

	//for {
	//	fmt.Println("abc")
	//}

	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	ch := make(chan string, 1) // Create a string channel

	// Goroutine sending data
	go func() {
		fmt.Println("sending hello")
		ch <- "hello go!!!"
		fmt.Println("sending 123")
		ch <- "123" // Sending data into the channel
	}()

	time.Sleep(5000 * time.Second)
	message := <-ch // Receiving data from the channel
	fmt.Println("Recived : ", message)
	message2 := <-ch
	fmt.Println("Recivedd: ", message2)
	message2 = ""
}
