package main

import "fmt"

type User struct {
	name string
	lastName string
	age int
}

/** Maps are unordered collections of key-value pairs */
func TypeMap() {
	fmt.Println(" ")
	fmt.Println("TYPES: Map")

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
		"name": "John",
		"lastName": "Doe",
	}
	usersMap := map[int]User{
		0: {
			name: "John",
			lastName: "Wick",
			age: 30,
		},
		1: {
			name: "Jane",
			lastName: "Doe",
			age: 25,
		},
	}
	fmt.Println("make person map:", makePersonMap)
	fmt.Println("make some map:", makeIntMap)
	fmt.Println("person map:", personMap)
	fmt.Println("users map:", usersMap)
	fmt.Println("--------------------------------")
}
