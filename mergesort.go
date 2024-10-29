package main

import (
	"fmt"
)

// merge function merges two sorted halves into a single sorted array.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge two sorted halves
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort function recursively divides the array and merges sorted halves
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point and split the array
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func main() {
	// Sample array to be sorted
	arr := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
