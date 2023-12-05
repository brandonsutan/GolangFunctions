package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declare a string variable
	var userInput string

	// Read input from the user
	fmt.Print("Enter a string: ")
	fmt.Scanln(&userInput)

	// Convert the string to lower case for case-insensitive comparison
	userInput = strings.ToLower(userInput)

	// Check if string contains 'i', 'a', and 'n' in that order
	if strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
