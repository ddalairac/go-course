package main

import (
	"fmt"
)

// Package scope variable only in this package
var pScope int = 42

// Global scope variable available in all packages
// (Definitions starting with capital letters are available everywhere)
var GScope int = 24

func BasicsTypes() {
	fmt.Println("TYPES basics")
	fmt.Println("package scope", pScope)
	fmt.Println("global scope", GScope)

	// define a type variable
	var whatToSay string = "Goodbye, cruel world"
	var i int = 7
	var f float32 = 3.14
	fmt.Println(whatToSay)
	fmt.Println("i:", i)
	fmt.Println("f:", f)

	// := infer type from return type
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
	fmt.Println("--------------------------------")
}

// functions
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
