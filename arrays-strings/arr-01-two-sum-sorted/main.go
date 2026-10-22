// Question: Given an integer slice sorted in nondecreasing order and a target, return
// 1-based indices of two distinct elements whose sum equals the target.
// Return an empty slice if no pair exists; any valid pair is acceptable.

package main

import "fmt"

const CONDITION = true

// Opposite-pointer scaffold and its sorted Two Sum implementation.
func fn(arr []int) int {
	// Two pointers: one input, opposite ends
	left := 0
	right := len(arr) - 1
	ans := 0

	for left <= right {
		// do some logic with left and right
		if CONDITION {
			left++
		} else {
			right--
		}
	}
	return ans
}

// twoSum returns 1-based indices, not Go's 0-based slice indices.
// time = O(n)
// space = O(1)
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
