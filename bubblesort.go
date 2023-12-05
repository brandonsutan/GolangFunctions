package main

import (
	"fmt"
)

func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	numbers := make([]int, 10)
	fmt.Println("Enter 10 integers:")

	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
		_, err := fmt.Scanf(&numbers[i])
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			i--
		}
	}

	BubbleSort(numbers)
	fmt.Println("Sorted slice:", numbers)
}
