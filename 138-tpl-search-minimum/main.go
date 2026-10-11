package main

import "fmt"

// minCapacity finds the smallest daily capacity that ships positive weights,
// in their original order, within days. days must be positive; total must fit int.
// Invariant: the minimum feasible capacity stays in [left,right]; right is feasible.
// Feasibility is monotonic: increasing capacity cannot require more days.
// Time O(n log(S+1)), S=sum(weights); space O(1).
func minCapacity(weights []int, days int) int {
	if days <= 0 {
		panic("days must be positive")
	}
	left, right := 0, 0
	for _, w := range weights {
		if w <= 0 {
			panic("weights must be positive")
		}
		left = max(left, w)
		right += w
	}
	feasible := func(capacity int) bool {
		used, load := 1, 0
		for _, w := range weights {
			if load > capacity-w {
				used++
				load = 0
			}
			load += w
		}
		return used <= days
	}
	for left < right {
		mid := left + (right-left)/2
		if feasible(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() { fmt.Println(minCapacity([]int{1, 2, 3, 4, 5}, 3)) }
