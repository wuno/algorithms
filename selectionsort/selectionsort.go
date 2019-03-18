package selectionsort

import "fmt"

// SelectionSort function takes an array of ints and returns a sorted array of ints
func SelectionSort(values []int) []int {

	fmt.Println("Selection Sort is Running...")
	fmt.Println("Selection Sort [Best/Worst: O(N^2)]")

	length := len(values)
	for i := 0; i < length; i++ {

		min := i
		for j := i; j < length; j++ {
			if values[j] < values[min] {
				min = j
			}
		}

		values[i], values[min] = values[min], values[i]

	}

	return values

}
