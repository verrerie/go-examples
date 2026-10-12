package main

import "fmt"

func main() {
	list := []int{1, 4, 2, 3, 4, 5, 6, 7, 6}
	fmt.Println(findPeakElement(list))
	list = []int{1, 2, 3}
	fmt.Println(findPeakElement(list))
	list = []int{1}
	fmt.Println(findPeakElement(list))
}

// LeetCode 162: Find Peak Element (requirements recalled from memory).
// A peak is strictly greater than its immediate neighbors. Treat the virtual
// boundary values nums[-1] and nums[len(nums)] as negative infinity.
// Return the INDEX of any peak, not necessarily the global maximum.
// For {1, 4, 2, 3, 4, 5, 6, 7, 6}, both indices 1 and 7 are valid answers.
// The solution must run in O(log n) time.
// Constraints:
//   - 1 <= len(nums) <= 1000
//   - -2^31 <= nums[i] <= 2^31 - 1
//   - nums[i] != nums[i+1] for every valid adjacent pair
//
// Time: O(log n). Space: O(1).
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// Keep an interval containing at least one peak; halve it each iteration.
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			// An upward slope guarantees a peak to the right, even if another
			// peak (possibly a taller one) is discarded on the left.
			l = mid + 1
		} else {
			// Adjacent values differ, so this is a downward slope.
			// A peak exists at mid or to its left.
			r = mid
		}
	}
	// l is now equal to r
	return l
}
