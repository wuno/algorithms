package bubblesort

import "fmt"

// BubbleSort function takes an array of ints and returns a sorted array of ints
func BubbleSort(values []int) []int {
	fmt.Println("Bubble Sort is Running...")
	fmt.Println("Bubble Sort [Best: O(N), Worst:O(N^2)]")

	l := len(values)

	// Loop the length of the collection passed in
	for i := 0; i < l; i++ {

		// Compare the next 2 values and swap them if the left value is greater
		for j := 0; j < (l - 1 - i); j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}

	}

	return values
}
