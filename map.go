package main

import "fmt"

func safeExecution() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	m := make(map[string]int)
	m["xyz"] = 2
	fmt.Println(m)
	fmt.Println(m["adsdsafds"])

	var myMap map[string]int // This map is nil
	myMap["test"] = 5        // ❌ Causes panic

	val, exist := m["pqr"]
	fmt.Println(val)
	fmt.Println(exist)
}

func main() {
	safeExecution() // If panic happens here, main() still runs after recovery

	// Execution continues
	fmt.Println("Program continues after panic!")

	str := new(string)
	*str = "124"
	fmt.Println(*str)
	str2 := "124"
	fmt.Println(str2 == *str)
}
