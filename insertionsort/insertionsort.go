package insertionsort

import "fmt"

// InsertionSort function takes an array of ints and returns a sorted array of ints
func InsertionSort(values []int) []int {

	fmt.Println("Insertion Sort is Running...")
	fmt.Println("Insertion Sort [Best: O(N), Worst:O(N^2)]")

	// Loop the length of the collection passed in
	// Starts on 1 so it can compare to the last node and move lowest to the left
	for i := 1; i < len(values); i++ {

		// Removes item and places it back in at the correct order placement
		j := i
		// If current index is less than last index, swap them
		for j > 0 && values[j] < values[j-1] {
			values[j], values[j-1] = values[j-1], values[j]
			// remove 1 from j till we get to 0
			// Move value backwards through collection to the correct placement
			// run insertion sort on the value iteratively
			j--
		}

	}

	return values
}
