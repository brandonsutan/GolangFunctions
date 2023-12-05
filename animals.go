package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal interface defines methods for the animal behaviors
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

// Cow, Bird, and Snake structs represent specific animal types
type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Methods for Cow
func (c Cow) Eat() string   { return c.food }
func (c Cow) Move() string  { return c.locomotion }
func (c Cow) Speak() string { return c.noise }

// Methods for Bird
func (b Bird) Eat() string   { return b.food }
func (b Bird) Move() string  { return b.locomotion }
func (b Bird) Speak() string { return b.noise }

// Methods for Snake
func (s Snake) Eat() string   { return s.food }
func (s Snake) Move() string  { return s.locomotion }
func (s Snake) Speak() string { return s.noise }

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput := scanner.Text()
		commands := strings.Split(userInput, " ")

		if len(commands) != 3 {
			fmt.Println("Invalid input. Please use the format: 'newanimal <name> <type>' or 'query <name> <info>'")
			continue
		}

		switch commands[0] {
		case "newanimal":
			animalName := commands[1]
			animalType := commands[2]
			switch animalType {
			case "cow":
				animals[animalName] = Cow{"grass", "walk", "moo"}
			case "bird":
				animals[animalName] = Bird{"worms", "fly", "peep"}
			case "snake":
				animals[animalName] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Invalid animal type. Available types are 'cow', 'bird', or 'snake'.")
			}
		case "query":
			animalName := commands[1]
			action := commands[2]
			animal, ok := animals[animalName]
			if !ok {
				fmt.Println("No such animal:", animalName)
				continue
			}
			switch action {
			case "eat":
				fmt.Println(animal.Eat())
			case "move":
				fmt.Println(animal.Move())
			case "speak":
				fmt.Println(animal.Speak())
			default:
				fmt.Println("Invalid action. Available actions are 'eat', 'move', or 'speak'.")
			}
		default:
			fmt.Println("Invalid command. Use 'newanimal' or 'query'.")
		}
	}
}
