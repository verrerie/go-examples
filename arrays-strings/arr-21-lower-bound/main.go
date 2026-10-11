package main

import "fmt"

// lowerBound returns the first index with value >= target, or n.
// Invariant: nums[:left] < target; nums[right:] >= target.
// The unknown interval [left,right) shrinks each step. Input is ascending.
// Time O(log n); space O(1).
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() { fmt.Println(lowerBound([]int{1, 3, 3, 5}, 3)) }
