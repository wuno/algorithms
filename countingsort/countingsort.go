package countingsort

import "fmt"

func getLargestNumber(values []int) int {

	largestNumber := 0
	for i := 0; i < len(values); i++ {
		if values[i] > largestNumber {
			largestNumber = values[i]
		}
	}

	return largestNumber

}

// CountingSort function takes an array of ints and returns a sorted array of ints
func CountingSort(values []int) []int {

	fmt.Println("Counting Sort is Running...")
	fmt.Println("Counting Sort [Best/Avg/Worst: O(N)]")

	// Find the largest value in the collection
	max := getLargestNumber(values) + 1

	// Innitialize a slice with the max number memory capacity and length defaults values to 0 in Go
	auxilory := make([]int, max)

	// Add the frequency of each number from original values collectiuon to the auxilory array index of that value
	for i := 0; i < len(values); i++ {
		auxilory[values[i]]++
	}

	// Add the frequency of each value to the index value index in auxilory
	for i, sum := 0, 0; i < max; i++ {
		sum, auxilory[i] = sum+auxilory[i], sum
	}

	// New array to hold our sorted values
	sortedValues := make([]int, len(values))
	for _, n := range values {
		sortedValues[auxilory[n]] = n
		auxilory[n]++
	}

	copy(values, sortedValues)

	return values

}
