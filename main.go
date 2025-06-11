package main

import (
	"fmt"
	"time"
)

// Package scope variable only in this package
var pScope int = 42

// Global scope variable available in all packages
// (Definitions starting with capital leters are available everyware)
var GScope int = 24

// type definition (for all packages because start with capital letter)
type Person struct {
	name string
	lastName string
	age  int
	birthday time.Time
}

func main() {
	fmt.Println("Hello, World!", pScope)

	// define a type variable
	var whatToSay string = "Goodbye, cruel world"

	fmt.Println(whatToSay)

	var i int = 7

	var f float32 = 3.14

	fmt.Println("i:", i)
	fmt.Println("f:", f)

	// := asign return type variable type
	whatWasSaid := saySomething(whatToSay)

	fmt.Println(whatWasSaid)

	// multiple return values
	whatWasSaid, whatElseWasSaid := sayTwothings(whatToSay)

	fmt.Println(whatWasSaid, whatElseWasSaid)

	// Pointer definition
	var myInt int = 42

	// & is used to get the address of the variable
	changeUsingPointer(&myInt)

	fmt.Println("myInt:", myInt)

	// struct
	person := Person{
		name: "John",
		lastName: "Doe",
		age: 30,
		birthday: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(person)
	fmt.Println("--------------------------------")
	// Structs()
}

func saySomething(whatToSay string) string {
	return whatToSay
}
func sayTwothings(whatToSay string) (string, string) {
	return whatToSay, "else"
}

func changeUsingPointer(z *int) {
	fmt.Println("z pointer:", z)
	*z = 24
}
