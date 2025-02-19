package main

import "fmt"

func testdefer() (result int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recpvered and modifying res")
			result = 10
		}
	}()
	panic("panic")
	return
}

func main() {

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recovered 1", r)
	//	}
	//}()

	fmt.Println("defer affected return: ", testdefer())
	var pr []int

	// Print statement before the panic
	fmt.Println("before risky call")
	func() {
		defer func() {
			if r := recover(); r == nil {
				fmt.Println("recovered inside risky unknown fn", r)
			}
		}()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered inside risky unknown fn", r)
			}
		}()

		// This will cause panic as pr is nil and we're trying to access index 2
		//pr[2] = 4
		fmt.Println("after errr inside fn")
	}()
	// Print statement after the panic (This won't be reached if panic is not recovered)
	fmt.Println("after risky call")

	// Code will continue after recovery
	fmt.Println("program continued execution after panic recovery")

	pr = make([]int, 2) // No more panic here
	fmt.Println(pr)

	arr := make([]int, 10)
	for x := 0; x <= 1; x++ {
		arr = append(arr, x)
	}

	// Further statements continue as normal
	fmt.Println("len of arr:", len(arr))
	fmt.Println("cap of arr:", cap(arr))

	// Dealing with a map
	mapp := make(map[string]int)
	mapp["apple"] = 5
	fmt.Println(len(mapp))

	m := map[string]int{"apple": 5, "banana": 10}
	for key, value := range m {
		if key == "apple" {
			delete(m, "apple")
		}
		fmt.Println(key, value)
	}
}
