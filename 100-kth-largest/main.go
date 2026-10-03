package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	list := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(list, 4))
}

// Expected time: O(n); worst case: O(n^2).
// Expected auxiliary space: O(log n); worst case: O(n) for the recursion stack.
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		panic("invalid input")
	}
	if len(nums) == 1 {
		return nums[0]
	}
	swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	pivot := nums[rand.IntN(len(nums))]
	l, r := 0, len(nums)-1
	i := 0
	// Partition time: O(m), where m is the current subarray length.
	for i <= r {
		switch {
		case nums[i] < pivot:
			swap(l, i)
			l++
			i++
		case nums[i] == pivot:
			i++
		case nums[i] > pivot:
			swap(r, i)
			r--
		}
	}
	smallCount := l
	bigCount := len(nums) - i
	equalCount := i - l
	switch {
	case k <= bigCount:
		return findKthLargest(nums[i:], k)
	case bigCount < k && k <= bigCount+equalCount:
		return pivot
	default:
		return findKthLargest(nums[:smallCount], k-bigCount-equalCount)
	}
}
