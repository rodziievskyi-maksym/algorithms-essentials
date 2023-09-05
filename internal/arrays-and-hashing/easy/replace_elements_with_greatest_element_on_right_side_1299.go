package arrays_and_hashing

import (
	"fmt"
)

func ReplaceElements() {
	/*
		Given an array arr, replace every element in that array with the greatest element among the elements to its right,
		and replace the last element with -1.

		After doing so, return the array.

		Input: arr = [17,18,5,4,6,1]
		Output: [18,6,6,6,1,-1]
		Explanation:
		- index 0 --> the greatest element to the right of index 0 is index 1 (18).
		- index 1 --> the greatest element to the right of index 1 is index 4 (6).
		- index 2 --> the greatest element to the right of index 2 is index 4 (6).
		- index 3 --> the greatest element to the right of index 3 is index 4 (6).
		- index 4 --> the greatest element to the right of index 4 is index 5 (1).
		- index 5 --> there are no elements to the right of index 5, so we put -1.

		Input: arr = [400]
		Output: [-1]
		Explanation: There are no elements to the right of index 0.
	*/

	result := replaceElements([]int{17, 18, 5, 4, 6, 1})
	fmt.Println("Replace Elements:", result)
}

func replaceElements(arr []int) []int {
	if len(arr) < 2 {
		return []int{-1}
	}

	// initial value is -1
	// reverse iteration -> We're going to start from the end of the array

	rightMax := -1
	// loop through the array in reverse
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > rightMax {
			arr[i], rightMax = rightMax, arr[i]
		} else {
			arr[i] = rightMax
		}
	}

	return arr
}
