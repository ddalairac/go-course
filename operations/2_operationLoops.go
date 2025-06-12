package operations

import (
	"fmt"
	"go-course/types"
)

/*
* Loops in Go
Key concepts:
- Go only has 'for' loops, but they can be used in different ways
- 'for' can be used as traditional for, while, and infinite loops
- 'range' is used to iterate over arrays, slices, maps, and strings
- Nested loops can be used for multi-dimensional iterations
*/
func OperationLoops() {
	fmt.Println(" ")
	fmt.Println("OPERATION: Loops")

	// Traditional for loop with initialization, condition, and increment
	for i := 0; i < 5; i++ {
		fmt.Println("for loop i:", i)
	}

	// While-like loop using only the condition
	// This is Go's way of implementing a while loop
	var j int = 0
	for j < 5 {
		fmt.Println("while loop j:", j)
		j++
	}

	// Infinite loop with break statement
	// Useful for continuous operations until a condition is met
	for {
		fmt.Println("infinite loop (with break statement)")
		break
	}

	// Nested loops for multi-dimensional iterations
	// Outer loop iterates 3 times, inner loop 3 times for each outer iteration
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("nested loop index:", i, "value:", j)
		}
	}

	// Range loop over a number
	// Iterates from 0 to n-1
	for i := range 5 {
		fmt.Println("range index:", i)
	}

	// Range loop over a slice
	// Returns index and value for each element
	sliceString := []string{"a", "b", "c", "d", "e"}
	for i, v := range sliceString {
		fmt.Println("range slice index:", i, "value:", v)
	}

	// Range loop over a map
	// Returns key and value for each element
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range myMap {
		fmt.Println("range map key:", k, "value:", v)
	}

	// Range loop over a string
	// Returns index and rune (Unicode character) for each character
	someString := "Hello, World!"
	fmt.Println("string:", someString)
	for i, v := range someString {
		fmt.Println("range string index:", i, "value:", v)
	}

	// Range loop over a slice of User struct
	users := []types.User{
		{Name: "John", LastName: "Doe", Age: 30},
		{Name: "Jane", LastName: "Smith", Age: 25},
		{Name: "Jim", LastName: "Brown", Age: 35},
	}
	for i, v := range users {
		fmt.Println("range slice of Users struct index:", i, "value:", v)
	}

	// Loop with continue statement
	// Skips the current iteration and continues with the next one
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("continue loop i:", i)
	}

	// Loop with break statement
	// Stops the loop immediately
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println("break loop i:", i)
	}

	// Loop with goto statement
	// Jumps to a labeled statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			goto label
		}
		fmt.Println("goto loop i:", i)
	}
label:
	fmt.Println("Jumped to label")
}
