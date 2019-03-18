package quicksort

import "fmt"

// QuickSort function takes an array of ints and returns a sorted array of ints
func QuickSort(values []int) []int {

	fmt.Println("Quick Sort is Running...")
	fmt.Println("Quick Sort [Best: O(N Lg N), Avg: O(N Lg N), Worst:O(N^2)]")

	if len(values) > 1 {

		// Create pivot point and 2 partitions
		pivot := len(values) / 2
		left := []int{}
		right := []int{}

		// Iterate over collection and sort by appending to partitions
		for i := range values {
			if i != pivot {
				if values[i] < values[pivot] {
					left = append(left, values[i])
				} else {
					right = append(right, values[i])
				}
			}
		}

		// Recursively call QuickSort on the partitions till they have been sorted
		QuickSort(left)
		QuickSort(right)

		// Combine partitioned array after QuickSort is performed
		var merged = append(append(append([]int{}, left...), []int{values[pivot]}...), right...)

		// Overwrite original array with current sorted array
		for j := 0; j < len(values); j++ {
			values[j] = merged[j]
		}

	}

	return values

}
