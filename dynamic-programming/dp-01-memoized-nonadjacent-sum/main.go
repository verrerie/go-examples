package main

import "fmt"

// maxNonAdjacent returns the largest sum of nonadjacent values; choosing none is allowed.
// State dp(i): best sum from nums[i:]. Sums must fit int.
// Invariant: each memo entry is the complete answer for that suffix.
// Time O(n); space O(n), including recursion. A zero answer is still cached.
func maxNonAdjacent(nums []int) int {
	memo := map[int]int{}
	var dp func(int) int
	dp = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if value, ok := memo[i]; ok {
			return value
		}
		memo[i] = max(dp(i+1), nums[i]+dp(i+2))
		return memo[i]
	}
	return dp(0)
}
func main() { fmt.Println(maxNonAdjacent([]int{2, 7, 9, 3, 1})) }
