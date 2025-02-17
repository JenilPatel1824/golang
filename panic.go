package main

import "fmt"

func riskyOperation() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered from panic:", r)
	//		panic("panic inside recover")
	//
	//	}
	//}()

	fmt.Println("Executing risky operation...")

	// This will cause a panic
	var myMap map[string]int // Nil map
	myMap["test"] = 5        // ❌ Causes panic
	panic("panic inside recover")

	fmt.Println("This will NOT execute due to panic.")
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after panic!") // ✅ Continues execution
}
