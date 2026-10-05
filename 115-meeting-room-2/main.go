package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minMeetingRooms([][]int{
		{0, 30},
		{5, 10},
		{15, 20},
		{16, 19},
	}))
}

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	myHeap := *h
	return myHeap[i] < myHeap[j]
}

func (h *MinHeap) Swap(i, j int) {
	myHeap := *h
	myHeap[i], myHeap[j] = myHeap[j], myHeap[i]
}

func (h *MinHeap) Push(x any) {
	i, _ := x.(int)
	*h = append(*h, i)
}

func (h *MinHeap) Pop() any {
	myHeap := *h
	top := myHeap[len(myHeap)-1]
	*h = myHeap[:len(myHeap)-1]
	return top
}

func (h *MinHeap) Peek() int {
	return (*h)[0]
}

// time = O(n * logn)
// space = O(n)
func minMeetingRooms(intervals [][]int) int {
	// time = O(n*logn)
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	h := make(MinHeap, 0)
	// time = O(n*logn), n iterations, and h per heap operation
	for _, interval := range intervals {
		if h.Len() == 0 {
			heap.Push(&h, interval[1])
		} else {
			if interval[0] >= h.Peek() {
				heap.Pop(&h)
			}
			heap.Push(&h, interval[1])
		}
	}
	return h.Len()
}
