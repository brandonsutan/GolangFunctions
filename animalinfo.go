package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct represents an animal with its behaviors
type Animal struct {
	eat   string
	move  string
	speak string
}

// Eat method returns the animal's eating behavior
func (a Animal) Eat() string {
	return a.eat
}

// Move method returns the animal's moving behavior
func (a Animal) Move() string {
	return a.move
}

// Speak method returns the animal's speaking behavior
func (a Animal) Speak() string {
	return a.speak
}

func main() {
	// Create a map of animals
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter requests in the format: animal action (e.g., 'cow speak'). Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Please use: animal action")
			continue
		}

		animal, ok := animals[parts[0]]
		if !ok {
			fmt.Println("Unknown animal. Please use 'cow', 'bird', or 'snake'.")
			continue
		}

		switch parts[1] {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			fmt.Println("Unknown action. Please use 'eat', 'move', or 'speak'.")
		}
	}
}
