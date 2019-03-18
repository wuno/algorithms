package heapsort

import "fmt"

func heapify(values []int, index int, size int) {

	// The purpose of heapify is to make the binary tree a max heap
	left := 2*index + 1
	right := left + 1

	var max int
	if left <= size && values[left] > values[index] {
		max = left
	} else {
		max = index
	}

	if right <= size && values[right] > values[max] {
		max = right
	}

	if max != index {
		top := values[index]
		values[index] = values[max]
		values[max] = top
		heapify(values, max, size)
	}

}

// HeapSort function takes an array of ints and returns a sorted array of ints
func HeapSort(values []int) []int {

	fmt.Println("Heap Sort is Running...")
	fmt.Println("Heap Sort [Best/Avg/Worst: O(N Lg N)]")

	length := len(values)
	mid := int(length / 2)

	// Make max heap so we can sort it
	for i := mid; i >= 0; i-- {
		heapify(values, i, length-1)
	}

	// Swap top node with bottomg node
	last := length - 1
	for j := last; j >= 0; j-- {
		top := values[0]
		values[0] = values[j]
		values[j] = top
		last--
		heapify(values, 0, last)
	}

	return values

}
