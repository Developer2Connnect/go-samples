package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Person represents the structure of the JSON data
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Read the JSON file
	fileContent, err := ioutil.ReadFile("example.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Create a Person struct to hold the data
	var person Person

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(fileContent, &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the result
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Email)
}
