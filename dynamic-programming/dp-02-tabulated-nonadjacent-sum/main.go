package main

import "fmt"

// maxNonAdjacent uses the same suffix state as the memoized version in dp-01-memoized-nonadjacent-sum.
// Invariant: before i, every dependency dp[i+1:] is already final.
// Fill backward because this recurrence depends on larger indices.
// Time O(n); space O(n). Choosing none is allowed; sums must fit int.
func maxNonAdjacent(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}
func main() { fmt.Println(maxNonAdjacent([]int{2, 7, 9, 3, 1})) }
