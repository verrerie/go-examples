package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(furthestBuilding(
		[]int{4, 12, 2, 7, 3, 18, 20, 3, 19},
		10,
		2,
	))
	fmt.Println(furthestBuilding(
		[]int{1, 2, 3, 4},
		0,
		0,
	))
	fmt.Println(furthestBuilding(
		[]int{4, 3, 2, 1},
		0,
		0,
	))
	fmt.Println(furthestBuilding(
		[]int{1, 1, 1, 1},
		0,
		0,
	))
}

// time = O(n*log(l+2)), l is the number of ladders
// space = O(l + 1)
func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := MinHeap{}
	for i := range len(heights) - 1 {
		climb := heights[i+1] - heights[i]
		if climb > 0 {
			if h.Len() < ladders {
				// O(logl)
				heap.Push(&h, climb)
			} else {
				// O(log(l+2))
				next := climb
				if h.Len() > 0 && h[0] < climb {
					next = h[0]
					h[0] = climb
					heap.Fix(&h, 0)
				}
				if bricks < next {
					return i
				} else {
					bricks = bricks - next
				}
			}
		}
	}
	return len(heights) - 1
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	last := len(old) - 1
	value := old[last]
	*h = old[:last]
	return value
}
