package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Structs implementing the interface
type Dog struct{}
type Cat struct{}

// Dog implements Speak() method
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat implements Speak() method
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts any Speaker
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {

	var a = 5
	fmt.Println(a)

	d := Dog{}
	c := Cat{}

	makeSound(d) // Woof!
	makeSound(c) // Meow!

}
