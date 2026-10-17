package main

import "fmt"

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3))
	fmt.Println(kWeakestRows([][]int{}, 1))
}

// time = O(mn + mlogk)
func kWeakestRows(mat [][]int, k int) []int {
	h := heap{
		[]position{},
		k,
	}
	// time = O(mn + mlogk)
	for r := range mat {
		strength := 0
		for _, n := range mat[r] {
			strength = strength + n
		}
		// 101 is higher than any row number, the formula combines the strength of the row and the row number
		// in order to simplify the comparison
		h.add(position{r, strength*101 + r})
	}
	ans := make([]int, k)
	// time = O(klogk)
	for h.len() > 0 {
		top := h.pop()
		ans[h.len()] = top.row
	}
	return ans
}

type position struct {
	row      int
	strength int
}

type heap struct {
	arr []position
	k   int
}

// time = O(logk)
func (h *heap) add(p position) {
	if h.len() < h.k {
		h.arr = append(h.arr, p)
		h.siftUp(h.len() - 1)
	} else if p.strength < h.peek().strength {
		h.arr[0] = p
		h.siftDown(0)
	}
	// otherwise just skip
}

// time = O(logk)
func (h *heap) siftUp(start int) {
	index := start
	for index > 0 {
		parent := (index - 1) / 2
		if h.arr[index].strength > h.arr[parent].strength {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

// time = O(logk)
func (h *heap) siftDown(start int) {
	index := start
	for index < h.len() {
		left := index*2 + 1
		right := index*2 + 2
		big := index
		if left < h.len() && h.arr[left].strength > h.arr[big].strength {
			big = left
		}
		if right < h.len() && h.arr[right].strength > h.arr[big].strength {
			big = right
		}
		if big == index {
			break
		}
		h.swap(big, index)
		index = big
	}
}

func (h *heap) swap(l, r int) {
	h.arr[l], h.arr[r] = h.arr[r], h.arr[l]
}

func (h *heap) peek() position {
	if h.len() > 0 {
		return h.arr[0]
	}
	return position{
		-1,
		-1,
	}
}

func (h *heap) len() int {
	return len(h.arr)
}

// time = O(k)
func (h *heap) pop() position {
	if h.len() == 0 {
		return position{-1, -1}
	}
	top := h.peek()
	h.swap(0, h.len()-1)
	h.arr = h.arr[:h.len()-1]
	h.siftDown(0)
	return top
}
