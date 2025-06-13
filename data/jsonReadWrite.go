// Package data provides functionality for handling JSON data operations
package data

import (
	"encoding/json"
	"fmt"
)

// PersonData represents a person's information with JSON tags for serialization
type PersonData struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}

// JsonReadWrite demonstrates reading from and writing to JSON in Go
// It shows how to:
// 1. Parse JSON string into Go structs
// 2. Convert Go structs back to JSON
func JsonReadWrite() {
	// Sample JSON data as a string containing an array of person objects
	jsonData := `[
		{
			"name": "John",
			"last_name": "Doe",
			"age": 30
		},
		{
			"name": "Jane",
			"last_name": "Smith",
			"age": 25
		}
	]`
	fmt.Println("jsonData", jsonData)
	fmt.Println(" ")

	// Convert JSON string to Go structs (Unmarshalling)
	unmarchalPersons := []PersonData{}
	err := json.Unmarshal([]byte(jsonData), &unmarchalPersons)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("unmarshalling JSON", unmarchalPersons)
	fmt.Println(" ")

	// Convert Go structs back to JSON string (Marshalling)
	jsonData2, err2 := json.Marshal(unmarchalPersons)
	if err2 != nil {
		fmt.Println("Error marshalling JSON:", err2)
	}
	fmt.Println("Marshal JSON", string(jsonData2))
}
