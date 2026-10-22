// Question: Given a nonempty n-by-n integer matrix whose rows and columns are
// sorted in nondecreasing order, return its kth smallest value, counting duplicate
// entries separately. k is one-based and satisfies 1 <= k <= n*n.

package main

import "fmt"

func main() {
	fmt.Println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8))
}

type point struct {
	v int
	c int
	r int
}

type heap []point

// time = O(n + klogn)
// space = O(n)
func kthSmallest(matrix [][]int, k int) int {
	// time = O(n)
	// space = O(n)
	h := heapifyFirstCols(matrix)
	n := len(h)
	// time = O(klogn)
	for k > 0 {
		p := pop(&h)
		k--
		if k == 0 {
			return p.v
		}
		newC := p.c + 1
		if newC == n {
			continue
		}
		add(&h, point{
			v: matrix[p.r][newC],
			c: newC,
			r: p.r,
		})
	}
	return -1
}

// after pop, the removed element is at L - 1
func pop(h *heap) point {
	hh := *h
	swap(hh, 0, len(hh)-1)
	ans := hh[len(hh)-1]
	hh = hh[:len(hh)-1]
	siftDown(hh, 0)
	*h = hh
	return ans
}

func add(h *heap, p point) {
	he := *h
	he = append(he, p)
	siftUp(he, len(he)-1)
	*h = he
}

func heapifyFirstCols(matrix [][]int) heap {
	h := make([]point, 0)
	for r := range len(matrix) {
		h = append(h, point{
			v: matrix[r][0],
			c: 0,
			r: r,
		})
	}
	for lastParent := len(h)/2 - 1; lastParent >= 0; lastParent-- {
		siftDown(h, lastParent)
	}
	return h
}

func swap(h heap, l, r int) {
	h[l], h[r] = h[r], h[l]
}

func siftDown(h heap, start int) {
	index := start
	for index < len(h) {
		left := index*2 + 1
		right := index*2 + 2
		small := index
		if left < len(h) && h[left].v < h[small].v {
			small = left
		}
		if right < len(h) && h[right].v < h[small].v {
			small = right
		}
		if small == index {
			break
		}
		swap(h, small, index)
		index = small
	}
}

func siftUp(h heap, start int) {
	index := start
	for index > 0 {
		parent := (index - 1) / 2
		if h[index].v < h[parent].v {
			swap(h, index, parent)
			index = parent
		} else {
			break
		}
	}
}
