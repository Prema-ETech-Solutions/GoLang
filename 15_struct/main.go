package main

import "fmt"

// struct
type Person struct {
	Name string
	Age  int
	Car  Car
}

// struct in struct
type Car struct {
	Brand string
	Year  int
}

// function to update values
func (person *Person) updatePerson(name string, age int, carBrand string, carYear int) {
	person.Name = name
	person.Age = age
	person.Car = Car{Brand: carBrand, Year: carYear}
}

func main() {
	// create instances of Person
	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Bob", Age: 25}
	fmt.Println(person1)
	fmt.Println(person2)

	// Accessing fields
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)

	// Update fields
	person1.Age = 31
	fmt.Println("Updated Age:", person1.Age)

	// Deleting fields
	person1 = Person{}
	fmt.Println("Deleted Person1:", person1)

	// Creating a new person
	person3 := Person{Name: "Charlie", Age: 28}
	fmt.Println("Created Person3:", person3)
	person3.updatePerson("Charlie Brown", 29, "Toyota", 2020)
	fmt.Println("Updated Person3:", person3)
}
