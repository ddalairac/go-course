package types

import (
	"fmt"
	"go-course/helpers"
)

type User struct {
	Name     string
	LastName string
	Age      int
}

/** Maps are unordered collections of key-value pairs */
func TypeMap() {
	helpers.PrintSubTitle("TYPES: Map")

	// Create an empty map using make() - useful when building the map dynamically
	makePersonMap := make(map[string]string)
	makePersonMap["name"] = "John"
	makePersonMap["lastName"] = "Wick"

	makeIntMap := make(map[int]int)
	makeIntMap[1] = 124323
	makeIntMap[2] = 22342
	makeIntMap[3] = 32342

	// Create and initialize a map using map literal syntax - when you know the elements upfront
	personMap := map[string]string{
		"name":     "John",
		"lastName": "Doe",
	}
	usersMap := map[int]User{
		0: {
			Name:     "John",
			LastName: "Wick",
			Age:      30,
		},
		1: {
			Name:     "Jane",
			LastName: "Doe",
			Age:      25,
		},
	}
	fmt.Println("make person map:", makePersonMap)
	fmt.Println("make some map:", makeIntMap)
	fmt.Println("person map:", personMap)
	fmt.Println("users map:", usersMap)
}
