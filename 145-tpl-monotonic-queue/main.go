package main

import "fmt"

// windowMax returns each size-k window's maximum. Invalid k returns empty.
// Invariant: deque indices increase, remain in the window, and values decrease.
// Remove expired indices at the front and dominated indices at the back.
// A ring buffer keeps O(k) state. Time O(n); output O(n-k+1).
func windowMax(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return []int{}
	}
	deque := make([]int, k)
	head, size := 0, 0
	out := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if size > 0 && deque[head] <= i-k {
			head = (head + 1) % k
			size--
		}
		for size > 0 && nums[deque[(head+size-1)%k]] <= v {
			size--
		}
		deque[(head+size)%k] = i
		size++
		if i >= k-1 {
			out = append(out, nums[deque[head]])
		}
	}
	return out
}
func main() { fmt.Println(windowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) }
