package mergesort

import "fmt"

func copyArray(values []int) []int {

	temp := make([]int, len(values))
	copy(temp, values)

	return temp

}

// MergeSort function takes an array of ints and returns a sorted array of ints
func MergeSort(values []int) []int {

	fmt.Println("Merge Sort is Running...")
	fmt.Println("Merge Sort [Best: n(log n), Worst: n(log n)]")

	if len(values) > 1 {
		mid := len(values) / 2
		left := copyArray(values[0:mid])
		right := copyArray(values[mid:])

		MergeSort(left)
		MergeSort(right)

		l := 0
		r := 0

		// Outer loop the total length of the collection
		for i := 0; i < len(values); i++ {

			lval := -1
			rval := -1

			if l < len(left) {
				lval = left[l]
			}

			if r < len(right) {
				rval = right[r]
			}

			if (lval != -1 && rval != -1 && lval < rval) || rval == -1 {
				values[i] = lval
				l++
			} else if (lval != -1 && rval != -1 && lval >= rval) || lval == -1 {
				values[i] = rval
				r++
			}

		}

	}

	return values

}
