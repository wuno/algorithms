package main

import (
	"fmt"

	bubble "github.com/wuno/algorithms/bubblesort"
	counting "github.com/wuno/algorithms/countingsort"
	heap "github.com/wuno/algorithms/heapsort"
	insertion "github.com/wuno/algorithms/insertionsort"
	merge "github.com/wuno/algorithms/mergesort"
	quick "github.com/wuno/algorithms/quicksort"
	radix "github.com/wuno/algorithms/radixsort"
	selection "github.com/wuno/algorithms/selectionsort"
)

var values = []int{34, 783, 12, 53, 9, 434, 123, 435, 43}

func main() {
	callInsertionSort()
}

func printSorted(values []int) {
	// Print the sorted array after sort function runs
	fmt.Printf("%v", values)
}

func callBubbleSort() {
	// Call BubbleSort on the array of values passed in
	printSorted(bubble.BubbleSort(values))
}

func callCountingSort() {
	// Call CountingSort on the array of values passed in
	printSorted(counting.CountingSort(values))
}

func callHeapSort() {
	// Call HeapSort on the array of values passed in
	printSorted(heap.HeapSort(values))
}

func callInsertionSort() {
	// Call InsertionSort on the array of values passed in
	printSorted(insertion.InsertionSort(values))
}

func callMergeSort() {
	// Call MergeSort on the array of values passed in
	printSorted(merge.MergeSort(values))
}

func callQuickSort() {
	// Call QuickSort on the array of values passed in
	printSorted(quick.QuickSort(values))
}

func callRadixSort() {
	// Call RadixSort on the array of values passed in
	printSorted(radix.RadixSort(values))
}

func callSelectionSort() {
	// Call SelectionSort on the array of values passed in
	printSorted(selection.SelectionSort(values))
}
