package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct that matches the desired JSON format
type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	// Create an instance of Person and populate it
	person := Person{
		Name:    "John Doe",
		Address: "123 Main St, Anytown, USA",
	}

	// Convert the Person instance to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}
