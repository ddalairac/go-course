package types

import (
	"fmt"

	"github.com/diegodalairac/go-course/app1_basics/helpers"
)

// Package scope variable only in this package
var pScope int = 42

// Global scope variable available in all packages
// (Definitions starting with capital letters are available everywhere)
var GScope int = 24

/** Basics primitive types */
func BasicsTypes() {
	helpers.PrintSubTitle("TYPES: Basics")
	fmt.Println("package scope", pScope)
	fmt.Println("global scope", GScope)

	// define a type variable
	var whatToSay string = "Goodbye, cruel world"
	var i int = 7
	var f float32 = 3.14
	var b bool = true
	fmt.Println(whatToSay)
	fmt.Println("i:", i)
	fmt.Println("f:", f)
	fmt.Println("b:", b)

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
