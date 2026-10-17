package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{1, 2, 3, 4, 5}))
	fmt.Println(lastStoneWeight([]int{1}))
	fmt.Println(lastStoneWeight([]int{1, 2, 3, 1000}))
	fmt.Println(lastStoneWeight([]int{1, 1}))
	fmt.Println(lastStoneWeight([]int{1, 1, 1, 1}))
	fmt.Println(lastStoneWeight(nil))
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

// time = O(n + logn * n)
// space = O(1)
func lastStoneWeight(stones []int) int {
	heapify(stones)
	for len(stones) > 1 {
		y := pop(&stones)
		x := pop(&stones)
		if x != y {
			add(&stones, y-x)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	} else {
		return 0
	}
}

// time = O(logn + amortized 1)
func add(arr *[]int, val int) {
	*arr = append(*arr, val)
	siftUp(*arr, len(*arr)-1)
}

// time = O(logn)
func siftUp(arr []int, start int) {
	index := start
	for index > 0 {
		parent := (index - 1) / 2
		if arr[index] > arr[parent] {
			swap(arr, index, parent)
			index = parent
		} else {
			break
		}
	}
}

// time = O(logn)
func pop(arr *[]int) int {
	top := (*arr)[0]
	swap(*arr, 0, len(*arr)-1)
	*arr = (*arr)[:len(*arr)-1]
	siftDown(*arr, 0)
	return top
}

// 0, 1, 2, 3, 4
// time = O(n)
func heapify(arr []int) {
	for lastParent := len(arr)/2 - 1; lastParent >= 0; lastParent-- {
		siftDown(arr, lastParent)
	}
}

// time = log(n)
func siftDown(arr []int, start int) {
	index := start
	for index < len(arr) {
		left := index*2 + 1
		right := index*2 + 2
		big := index
		if left < len(arr) && arr[left] > arr[big] {
			big = left
		}
		if right < len(arr) && arr[right] > arr[big] {
			big = right
		}
		if big == index {
			break
		}
		swap(arr, big, index)
		index = big
	}
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}
