package data

import (
	"encoding/json"
	"fmt"
)

type PersonData struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}

func JsonReadWrite() {
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

	// Read json data
	unmarchalPersons := []PersonData{}
	err := json.Unmarshal([]byte(jsonData), &unmarchalPersons)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	fmt.Println("unmarshalling JSON", unmarchalPersons)
	fmt.Println(" ")

	// Write json data
	jsonData2, err2 := json.Marshal(unmarchalPersons)
	if err2 != nil {
		fmt.Println("Error marshalling JSON:", err2)
	}
	fmt.Println("Marshal JSON", string(jsonData2))
}
