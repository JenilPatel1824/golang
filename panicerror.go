package main

import "fmt"

func main() {
	// Defer a function that recovers from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println(10 / 0)
	fmt.Println("kndee" + "ehv2e" + 1)

	// This will cause a panic, but it will be recovered
	//const a int = 0
	//result := 10 / a
	//fmt.Println(result)

	fmt.Println(10 / 0)

}
