package main

import (
	"fmt"
	"strings"
)

func main() {

	//var arr2 [4]string
	//arr2[0] = "xyz"
	arr := [4]string{"a", "b", "c", "d"}
	arr2 := [4]string{"a", "b", "c", "d"}
	fmt.Println(arr == arr2)

	arr3 := [3]int{1, 2, 3}
	arr4 := [3]int{1, 2, 3}

	var destination *[4]string = &arr

	for i, j := range arr {
		fmt.Printf("index: %d val %s", i, j)
		fmt.Println()
	}

	fmt.Println(*destination)
	destination[0] = "pqr"
	fmt.Println(arr)

	slice := make([]int, 10, 100)

	var slice2 [10]int
	fmt.Println(arr3 == arr4)
	fmt.Println(slice2[0])
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	s := "@@Hello, Geeks"
	result := strings.TrimRight(s, "s")

	fmt.Println(result)

	splittest := "hello,123,453,world"
	splitarr := strings.SplitAfter(splittest, ",")
	fmt.Println(splitarr)
	fmt.Println(string(splitarr[0][0]))

	fmt.Println("abcd" < "abc")

	var x = 100
	var p = &x

	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)
}
