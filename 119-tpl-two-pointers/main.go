package main

import "fmt"

func main() {
	fmt.Println(fn([]int{1, 2, 3, 4}))
	fmt.Println(fn2([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(fn3([]int{1, 3, 5, 7}))
}

const CONDITION = true

func fn(arr []int) int {
	// Two pointers: one input, opposite ends
	left := 0
	right := len(arr) - 1
	ans := 0

	for left < right {
		// do some logic with left and right
		if CONDITION {
			left++
		} else {
			right--
		}
	}
	return ans
}

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

const WINDOW_CONDITION_BROKEN = false

func fn3(arr []int) int {
	var left, ans, curr int

	for right := range len(arr) {
		// do logic to add arr[right] to curr
		curr = curr + arr[right]
		for WINDOW_CONDITION_BROKEN {
			// remove arr[left] from curr
			curr = curr - arr[left]
			left++
		}

		// update ans
		ans = ans + curr
	}

	return ans
}
