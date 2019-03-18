package radixsort

import (
	"fmt"
	"strconv"
)

func getLargestNumber(values []int) int {

	largestNumber := 0
	for i := 0; i < len(values); i++ {
		if values[i] > largestNumber {
			largestNumber = values[i]
		}
	}

	return largestNumber

}

// RadixSort function takes an array of ints and returns a sorted array of ints
func RadixSort(values []int) []int {

	fmt.Println("Radix Sort is Running...")
	fmt.Println("Radix Sort [Best/Avg/Worst: O(N)]")

	largestNumber := getLargestNumber(values)
	size := len(values)
	significantDigit := 1
	auxilory := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for largestNumber/significantDigit > 0 {

		fmt.Println("\tSorting: "+strconv.Itoa(significantDigit)+"'s place", values)

		bucket := [10]int{0}

		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(values[i]/significantDigit)%10]++
		}

		// Add the count of the previous buckets
		// Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the bucket to fill a "auxilory" array
		for i := size - 1; i >= 0; i-- {
			bucket[(values[i]/significantDigit)%10]--
			auxilory[bucket[(values[i]/significantDigit)%10]] = values[i]
		}

		// Replace the current array with the auxilory array
		for i := 0; i < size; i++ {
			values[i] = auxilory[i]
		}

		fmt.Println("\n\tBucket: ", bucket)

		// Move to next significant digit
		significantDigit *= 10
	}

	return values

}
