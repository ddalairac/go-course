package types

import (
	"fmt"
	"go-course/helpers"
	"sort"
)

/*
* Slices are dynamic arrays that can grow or shrink as needed
Key differences with arrays:
- Arrays have fixed size, slices are dynamic
- Arrays: [5]int{} - size is part of the type
- Slices: []int{} - size can change
- Slices can be created from arrays using slicing
- Slices can be sorted and modified
*/
func TypeSlice() {
	helpers.PrintSubTitle("TYPES: Slice")

	// Create a slice using make() with initial length 3 and capacity 5
	// This is useful when you know you'll need a certain capacity
	makeSlice := make([]int, 3, 5)
	makeSlice[0] = 1
	makeSlice[1] = 2
	makeSlice[2] = 3

	// Create a slice using slice literal syntax
	// This is simpler when you know the initial elements
	sliceInt := []int{4, 5, 6}

	// Append items to slice - this is not possible with arrays
	// The slice will grow automatically if needed
	sliceInt = append(sliceInt, 1, 2, 3)

	// Create a slice from an existing array
	// This creates a view into the array from index 1 to 3 (exclusive)
	array := [5]int{1, 2, 3, 4, 5}
	sliceFromArray := array[1:3]

	fmt.Println("make slice:", makeSlice)
	fmt.Println("slice:", sliceInt)
	sort.Ints(sliceInt)
	fmt.Println("sorted slice:", sliceInt)
	fmt.Println("slice from array:", sliceFromArray)
}
