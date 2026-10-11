package main

import "fmt"

// Opposite-pointer scaffold and its sorted Two Sum implementation.
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
