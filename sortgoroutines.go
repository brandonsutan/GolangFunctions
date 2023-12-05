package main

import (
	"fmt"
	"sort"
	"sync"
)

// Function to sort a sub-array
func sortSubArray(arr []int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine sorting sub-array:", arr)
	sort.Ints(arr)
	wg.Done()
}

// Main function
func main() {
	// Example array to sort
	arr := []int{9, 3, 6, 2, 10, 5, 8, 4, 7, 1}
	n := len(arr)
	subArraySize := n / 4

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create slices to hold sorted sub-arrays
	sortedSubArrays := make([][]int, 4)

	// Start four goroutines to sort sub-arrays
	wg.Add(4)
	for i := 0; i < 4; i++ {
		from := i * subArraySize
		to := from + subArraySize
		if i == 3 { // Ensure the last goroutine sorts any remaining elements
			to = n
		}

		subArray := make([]int, to-from)
		copy(subArray, arr[from:to])
		sortedSubArrays[i] = subArray

		go sortSubArray(subArray, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Merge the sorted sub-arrays
	mergedArray := make([]int, 0, n)
	for _, subArr := range sortedSubArrays {
		mergedArray = append(mergedArray, subArr...)
	}

	// Sort the merged array to complete the merge process
	sort.Ints(mergedArray)

	fmt.Println("Merged and sorted array:", mergedArray)
}
