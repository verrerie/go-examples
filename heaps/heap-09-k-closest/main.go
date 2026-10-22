// Question: Given integer points [x, y] and 1 <= k <= len(points), return k
// points closest to the origin by Euclidean distance, in any order. Any selection
// among equally distant points is valid. Assume squared distances fit int.

package main

import "fmt"

func main() {
	fmt.Println(kClosest([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
	fmt.Println(kClosest([][]int{{1, 1}}, 1))
	fmt.Println(kClosest([][]int{
		{2, 2}, {-2, -2}, {2, -2}, {-2, 2},
	}, 3))
	fmt.Println(kClosest([][]int{
		{1, 0}, {2, 0}, {3, 0}, {4, 0}, {0, 0},
	}, 4))
}

// time = O(nlog(k+1) + k)
// space = O(k)
func kClosest(points [][]int, k int) [][]int {
	// space = O(k)
	h := NewHeap(k)
	// time = O(nlog(k+1))
	for _, coord := range points {
		h.add(point{coord, coord[0]*coord[0] + coord[1]*coord[1]})
	}
	// excluding result as space overhead
	ans := make([][]int, k)
	// time = O(k)
	for i, p := range h.arr {
		ans[i] = p.coord
	}
	return ans
}

type point struct {
	coord []int
	dist  int
}

type heap struct {
	arr  []point
	size int
	k    int
}

func NewHeap(k int) heap {
	return heap{make([]point, k), 0, k}
}

// invariant: store closest points in max-heap, worst point is on top
// time = O(logk), space = O(1)
func (h *heap) add(p point) {
	if h.size < h.k {
		h.arr[h.size] = p
		h.size++
		h.siftUp(h.size - 1)
	} else if p.dist < h.arr[0].dist {
		h.arr[0] = p
		h.siftDown(0)
	}
}

// time = O(logk)
func (h *heap) siftDown(start int) {
	index := start
	for index < h.size {
		left := index*2 + 1
		right := index*2 + 2
		big := index
		if left < h.size && h.arr[left].dist > h.arr[big].dist {
			big = left
		}
		if right < h.size && h.arr[right].dist > h.arr[big].dist {
			big = right
		}
		if index == big {
			break
		}
		h.swap(index, big)
		index = big
	}
}

func (h *heap) swap(l, r int) {
	h.arr[l], h.arr[r] = h.arr[r], h.arr[l]
}

// time = O(logk)
func (h *heap) siftUp(start int) {
	index := start
	for index > 0 {
		parent := (index - 1) / 2 // test: 0,1,2,3,4,5, correct
		if h.arr[index].dist > h.arr[parent].dist {
			h.swap(parent, index)
			index = parent
		} else {
			break
		}
	}
}
