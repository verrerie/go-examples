package main

import "fmt"

func main() {
	fmt.Println(connectSticks([]int{1, 8, 3, 5}))
	fmt.Println(connectSticks([]int{10}))
	fmt.Println(connectSticks([]int{}))
	fmt.Println(connectSticks([]int{1, 2, 3, 4, 5}))
	fmt.Println(connectSticks([]int{2, 2, 2}))
}

func connectSticks(sticks []int) int {
	heapify(sticks)
	cost := 0
	for len(sticks) > 1 {
		first := pop(&sticks)
		second := pop(&sticks)
		combined := first + second
		cost = cost + combined
		add(&sticks, combined)
	}
	return cost
}

func add(arr *[]int, e int) {
	heap := *arr
	heap = append(heap, e)
	siftUp(heap, len(heap)-1)
	*arr = heap
}

func siftUp(arr []int, current int) {
	for current > 0 {
		parent := (current - 1) / 2
		if arr[current] < arr[parent] {
			swap(arr, current, parent)
			current = parent
		} else {
			break
		}
	}
}

func pop(arr *[]int) int {
	heap := *arr
	top := heap[0]
	swap(heap, 0, len(heap)-1)
	heap = heap[:len(heap)-1]
	siftDown(heap, 0)
	*arr = heap
	return top
}

func heapify(arr []int) {
	for parent := len(arr)/2 - 1; parent >= 0; parent-- {
		siftDown(arr, parent)
	}
}

func siftDown(arr []int, current int) {
	for current < len(arr) {
		left := current*2 + 1
		right := current*2 + 2
		smaller := current
		if left < len(arr) && arr[left] < arr[smaller] {
			smaller = left
		}
		if right < len(arr) && arr[right] < arr[smaller] {
			smaller = right
		}
		if smaller == current {
			break
		}
		swap(arr, smaller, current)
		current = smaller
	}
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}
