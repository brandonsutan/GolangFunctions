package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	slice := make([]int, 0)

	for {
		fmt.Println("Enter an integer:")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			continue
		}

		slice = append(slice, n)
		sort.Ints(slice)

		fmt.Println("Sorted slice:", slice)
	}
}
