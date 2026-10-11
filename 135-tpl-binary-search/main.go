package main

import "fmt"

// search returns any matching index, or -1, in an ascending array.
// Invariant: if a match exists, at least one remains inside [left,right].
// Time O(log n); space O(1).
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // At this point left would be the insertion index for an absent target.
}
func main() { fmt.Println(search([]int{1, 3, 5, 7}, 5)) }
