package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if a filename is provided
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file name")
	}
	fileName := os.Args[1]

	// Read the file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Split the file content into lines
	lines := strings.Split(string(content), "\n")

	// Iterate over lines and print first name/last name pairs
	for _, line := range lines {
		if line != "" {
			nameParts := strings.Fields(line)
			if len(nameParts) >= 2 {
				fmt.Printf("First Name: %s, Last Name: %s\n", nameParts[0], nameParts[1])
			}
		}
	}
}
