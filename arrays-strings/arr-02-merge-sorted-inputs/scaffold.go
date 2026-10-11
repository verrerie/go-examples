package main

// Original two-input teaching scaffold. mergeSorted is the runnable implementation.
const CONDITION = true

func fn2(arr1, arr2 []int) int {
	// Two pointers: two inputs, exhaust both
	i, j, ans := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		// do logic for i and j
		if CONDITION {
			i++
		} else {
			j++
		}
	}

	for i < len(arr1) {
		// do logic for i
		i++
	}
	for j < len(arr2) {
		// do logic for j
		j++
	}
	return ans
}
