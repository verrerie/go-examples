package main

import "fmt"

// countSum counts nonempty subarrays with sum target. Negative values are allowed.
// Invariant: counts stores frequencies of prefixes strictly before this endpoint.
// Query before insertion, so an empty subarray is never counted.
// Time O(n) expected; space O(n). Sums and counts must fit int.
func countSum(nums []int, target int) int {
	counts := map[int]int{0: 1} // The empty prefix allows ranges starting at 0.
	sum, answer := 0, 0
	for _, v := range nums {
		sum += v
		answer += counts[sum-target]
		counts[sum]++
	}
	return answer
}
func main() { fmt.Println(countSum([]int{1, -1, 0}, 0)) }
