// Question: Given a nonempty integer slice and 1 <= k <= len(nums), return the kth
// largest value, counting duplicate occurrences separately. k is 1-based.

package main

import (
	"container/heap"
	"fmt"
	"math/rand/v2"
)

func main() {
	list := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(list, 4))
	fmt.Println(findKthLargestHeap(list, 4))
}

// findKthLargestHeap returns the kth largest value without changing nums.
// Invariant: after processing at least k values, the min-heap contains the k
// largest values seen, counting duplicates. Its root is the smallest retained
// value, so it is the kth largest of the processed values.
// Time: O(k + (n-k)*log(k+1)): heapify k values, then scan the remaining n-k.
// Each scan compares in O(1); only a larger value needs an O(log(k+1)) root repair.
// Extra space: O(k) for the copied heap. Invalid k panics before indexing nums.
func findKthLargestHeap(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		panic("k out of range")
	}
	h := minHeap(append([]int(nil), nums[:k]...))
	heap.Init(&h)
	for _, value := range nums[k:] {
		if value > h[0] {
			h[0] = value
			heap.Fix(&h, 0)
		}
	}
	return h[0]
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop operate on the last slot; container/heap handles reordering.
func (h *minHeap) Push(value any) { *h = append(*h, value.(int)) }
func (h *minHeap) Pop() any {
	last := len(*h) - 1
	value := (*h)[last]
	*h = (*h)[:last]
	return value
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
