package main

import (
	"fmt"
	"time"
)

/** type definition (for all packages because start with capital letter)*/
type Person struct {
	name string
	lastName string
	age  int
	birthday time.Time
}

// struct method definition
func (instance *Person) fullName() string {
	return instance.name + " " + instance.lastName
}

// Structs are collections of fields
func StructsTypes() {
	fmt.Println(" ")
	fmt.Println("TYPES: Structs")

	// Create a new Person struct with specific values
	person := Person{
		name: "John",
		lastName: "Wick",
		age: 40,
		birthday: time.Date(1985, 10, 19, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println("person:", person)

	// Initialize a Person struct with default zero values
	var otherPerson Person
	fmt.Println("otherPerson:", otherPerson)

	// Call the struct method to get the full name
	fmt.Println("person fullname:", person.fullName())

	fmt.Println("--------------------------------")
}
