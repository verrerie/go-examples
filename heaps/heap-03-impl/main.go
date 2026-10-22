// Question: Implement an integer min-heap supporting insertion, reading the
// minimum, and removing and returning the minimum. Allow duplicate values and
// growth beyond the initial capacity; peek and pop on an empty heap must panic.

package main

import "fmt"

func main() {
	h := NewHeap(10)
	for i := range 10 {
		h.add(10 - i)
	}
	fmt.Println(h.minHeap, h.size)
	for range 10 {
		fmt.Println(h.pop())
	}
	h.add(1)
	h.add(1)
	h.add(1)
	h.add(0)
	fmt.Println(h.pop(), h.pop(), h.pop(), h.pop())
}

type Heap struct {
	minHeap []int
	size    int
}

func NewHeap(size int) *Heap {
	return &Heap{
		// we use index start by 1 to have simpler parent - child index calculations
		// root index is 1
		minHeap: make([]int, 1, size+1),
		size:    0,
	}
}

func (h *Heap) add(e int) {
	// adding algo:
	// 1. append to the last child
	// 2. check its parent
	// 3. iterate on parent of parent
	h.size++
	h.minHeap = append(h.minHeap, e) // amortized O(1)

	// space = O(1)
	index := h.size
	parent := index / 2
	// time = O(h) = O(logn), every step moves 1 level upward
	for index > 1 && h.minHeap[index] < h.minHeap[parent] {
		swap(h.minHeap, index, parent)
		index = parent
		parent = index / 2
	}
}

func (h *Heap) peek() int {
	if h.size < 1 {
		panic("empty heap")
	}
	return h.minHeap[1]
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func (h *Heap) pop() int {
	if h.size < 1 {
		panic("no more element")
	}
	// space = O(1)
	ans := h.peek()
	swap(h.minHeap, 1, h.size)
	h.size--
	h.minHeap = h.minHeap[:h.size+1]

	index := 1

	// time = O(h) = O(logn), every iteration move one level downward
	for index <= h.size/2 {
		left := index * 2
		right := index*2 + 1

		smallest := left
		if right <= h.size && h.minHeap[right] < h.minHeap[left] {
			smallest = right
		}

		if h.minHeap[smallest] < h.minHeap[index] {
			swap(h.minHeap, index, smallest)
			index = smallest
		} else {
			// exit early
			break
		}
	}
	return ans
}
