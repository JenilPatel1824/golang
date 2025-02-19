package main

import (
	json2 "encoding/json"
	"fmt"
)

var pqr int

type Person struct {
	name string
	age  int
}

func (p Person) Greet() {
	p.name = "new name"
	fmt.Println("Hello, my name is", p.name)
}

// Struct Person has methods on both value and pointer receivers. Such usage is not recommended by the Go
func (p *Person) Birthday() {
	p.age++
}
func main() {

	d := 10
	defer fmt.Println(d)
	d = 10
	p := &Person{"Alice", 25}
	p.age = 40
	// Order must match field declaration
	fmt.Println(*p)
	p.Greet()
	fmt.Println(p)
	// Output: {Alice 25}
	p2 := Person{name: "Alice", age: 25}
	fmt.Println(p2)
	p3 := new(Person) // Returns a pointer *Person
	p3.name = "Charlie"
	p3.age = 40
	fmt.Println(*p3)
	fmt.Println("json of p3")
	json, other := json2.Marshal(*p3)
	fmt.Println(string(json))

	fmt.Println(other)
	p4 := Person{}
	fmt.Println(p4)
	p5 := Person{name: "John", age: 30}
	fmt.Println(p5)
	// Output: {Charlie 40}
	// Output: {Alice 25}

	//ch := make(chan int) // Create a channel

	var x int
	fmt.Println(x)
	var y = 5
	fmt.Println(y)

	// This will block because there is no sender
	//val := <-ch
	//close(ch)
	fmt.Print("trying to recive: ")
	//fmt.Println(<-ch)
	//val, open := <-ch
	//fmt.Println(val, open)

	fmt.Println() //ch <- fmt.Println("Received:")
	// main func2tion ends without sending, receiver will be blocked forever

	s := [4]int{1, 23, 4}
	fmt.Println(s)

	//var a = "Java"
	//var b = "Guides"
	fmt.Println("a", 5)
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	var m1, m2 map[string]int
	fmt.Println(m1 == nil) // true
	fmt.Println(m2 == nil)

	const (
		pr = 7
		q
		z
	)
	fmt.Println(pr, q, z)
	str := "xyz"
	fmt.Println("below: ")
	fmt.Println(string(str[4]))
}
