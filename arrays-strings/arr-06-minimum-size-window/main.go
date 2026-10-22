// Question: Given a positive target and a slice of positive integers, return the
// minimum length of a contiguous subarray with sum at least target.
// Return 0 if no such subarray exists.

package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1}))
}

// time = O(n)
// time = O(1)
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	minL := 0

	sum := 0
	for left < len(nums) && right < len(nums) {
		for sum < target && right < len(nums) {
			sum = sum + nums[right]
			right++
		}
		for sum >= target {
			if minL == 0 {
				minL = right - left
			} else {
				minL = min(minL, right-left)
			}
			sum = sum - nums[left]
			left++
		}
	}
	return minL
}
