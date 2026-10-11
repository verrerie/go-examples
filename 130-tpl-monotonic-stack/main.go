package main

import "fmt"

// nextSmaller returns the next strictly smaller item's index, or -1.
// Invariant: stack indices increase; their values do not decrease.
// Each stacked index still needs its first smaller item to the right.
// Time O(n): each index is pushed and popped at most once. Space O(n).
func nextSmaller(nums []int) []int {
	out := make([]int, len(nums))
	for i := range out {
		out[i] = -1
	}
	stack := []int{}
	for i, v := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > v {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			out[j] = i
		}
		stack = append(stack, i)
	}
	return out
}

func main() {
	fmt.Println(nextSmaller([]int{3, 1, 2, 1}))
	fmt.Println(nextSmaller([]int{3, 4, 5, 6, 1, 0}))
}
