package main

import "fmt"

// interface
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() string {
	return "Meow! My name is " + c.Name
}

func main() {
	var animal Animal
	animal = &Dog{Name: "Buddy"}
	fmt.Println(animal.Speak())
	animal = &Cat{Name: "Whiskers"}
	fmt.Println(animal.Speak())
}
