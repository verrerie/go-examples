package main

import "fmt"

// maxPiece finds the largest integer piece length yielding at least k pieces.
// Lengths are nonnegative and k is positive. Zero means no positive solution.
// Invariant: the maximum feasible length stays in [left,right]; 0 is a sentinel.
// Feasibility is monotonic: smaller positive pieces cannot reduce the count.
// Time O(n log(M+1)), M=max(lengths); space O(1).
func maxPiece(lengths []int, k int) int {
	if k <= 0 {
		panic("k must be positive")
	}
	left, right := 0, 0
	for _, v := range lengths {
		if v < 0 {
			panic("negative length")
		}
		right = max(right, v)
	}
	feasible := func(size int) bool {
		remaining := k
		for _, v := range lengths {
			pieces := v / size
			if pieces >= remaining {
				return true
			}
			remaining -= pieces
		}
		return false
	}
	for left < right {
		delta := right - left
		mid := left + delta/2 + delta%2 // Upper midpoint guarantees progress.
		if feasible(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
func main() { fmt.Println(maxPiece([]int{7, 5, 9}, 5)) }
