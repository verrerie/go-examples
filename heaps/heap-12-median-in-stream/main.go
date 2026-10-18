package main

import (
	"container/heap"
	"fmt"
)

func main() {
	m := Constructor()
	m.AddNum(1)
	fmt.Println(m.FindMedian())
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	m.AddNum(-4)
	fmt.Println(m.FindMedian())
}

type Heap struct {
	arr []int
	min bool
}

func (h *Heap) Len() int {
	return len(h.arr)
}

func (h *Heap) Less(i, j int) bool {
	if h.min {
		return h.arr[i] < h.arr[j]
	}
	return h.arr[i] > h.arr[j]
}

func (h Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *Heap) Pop() any {
	last := h.Len() - 1
	value := h.arr[last]
	h.arr = h.arr[:last]
	return value
}

func (h *Heap) Peek() int {
	return h.arr[0]
}

// MedianFinder uses O(n) space after n insertions.
type MedianFinder struct {
	lower  Heap
	higher Heap
}

// Constructor uses O(1) time and space.
func Constructor() MedianFinder {
	return MedianFinder{
		lower:  Heap{[]int{}, false},
		higher: Heap{[]int{}, true},
	}
}

// AddNum time = O(logn)
func (m *MedianFinder) AddNum(num int) {
	if m.lower.Len() == 0 || num < m.lower.Peek() {
		// O(log(n/2)
		heap.Push(&m.lower, num)
		if m.lower.Len()-m.higher.Len() > 1 {
			// O(log(n/2)
			val := heap.Pop(&m.lower)
			// O(log(n/2)
			heap.Push(&m.higher, val)
		}
	} else {
		heap.Push(&m.higher, num)
		if m.higher.Len()-m.lower.Len() > 0 {
			val := heap.Pop(&m.higher)
			heap.Push(&m.lower, val)
		}
	}
}

// FindMedian time = O(1)
func (m *MedianFinder) FindMedian() float64 {
	// odd case
	if m.lower.Len() > m.higher.Len() {
		return float64(m.lower.Peek())
	} else {
		// even case
		return float64(m.lower.Peek()+m.higher.Peek()) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
