package operations

import (
	"fmt"

	"github.com/diegodalairac/go-course/app1_basics/helpers"
)

/*
* Operation and conditionals
Key concepts:
- if/else statements work similar to other languages
- else if can be chained multiple times
- switch statements can be used for multiple conditions
- switch cases don't need break statements (unlike other languages)
*/
func OperationCondition() {
	helpers.PrintSubTitle("OPERATION: Condition")

	var isTrue bool = true

	// Basic if/else if/else structure
	// Note: Go doesn't use parentheses for conditions
	if isTrue {
		fmt.Println("true")
	} else if !isTrue { // Logical NOT operator (!) works like JS
		fmt.Println("false")
	} else {
		fmt.Println("unknown")
	}

	var intVar int = 1
	// Switch statement with multiple cases
	// Unlike other languages, Go automatically breaks after each case
	// No need for explicit break statements
	switch intVar {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unknown")
	}

}
