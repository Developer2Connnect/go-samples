// main.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Person represents the structure of the YAML data
type Person struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Email string `yaml:"email"`
}

func main() {
	// Read the YAML file
	fileContent, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Create a Person struct to hold the data
	var person Person

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(fileContent, &person)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Print the result
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Email)
}
