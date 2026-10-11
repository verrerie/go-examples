package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

// topK returns the largest k values in descending order; duplicates count.
// Invariant: the min-heap holds the largest min(k, processed) values seen.
// For k<=0 return empty; for k>n keep all n. Input is unchanged.
// Time O(n log(k+1)); space O(k), after clamping k to [0,n].
func topK(nums []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	k = min(k, len(nums))
	h := &minHeap{}
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	out := make([]int, h.Len())
	for i := len(out) - 1; i >= 0; i-- {
		out[i] = heap.Pop(h).(int)
	}
	return out
}
func main() { fmt.Println(topK([]int{3, 1, 5, 5, 2}, 3)) }
