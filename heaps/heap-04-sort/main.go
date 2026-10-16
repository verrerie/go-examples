package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7}
	maxHeapify(l)
	fmt.Println(l)
	l = []int{0}
	maxHeapify(l)
	fmt.Println(l)
	l = []int{1, 2, 3, 4}
	maxHeapify(l)
	fmt.Println(l)
	l = []int{1, 2, 3}
	maxHeapify(l)
	fmt.Println(l)
	l = []int{1, 2, 3, 4, 5}
	heapSort(l)
	fmt.Println(l)
	l = []int{1, 2, 2, 3, 3}
	heapSort(l)
	fmt.Println(l)
	l = []int{5, 4, 3, 2, 1}
	heapSort(l)
	fmt.Println(l)
	l = []int{1}
	heapSort(l)
	fmt.Println(l)
	l = []int{7, 2, 4, 1, 6, 0}
	heapSort(l)
	fmt.Println(l)
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func maxHeapify(arr []int) {
	// space = O(1)
	// time = O(n) = O(n/4 × 1 + n/8 × 2 + n/16 × 3 + …)
	for lastParent := len(arr)/2 - 1; lastParent >= 0; lastParent-- {
		// eg. len = 5, lastParent = 1 (2,3,4 are leaves), left  = 3, right = 4
		siftDown(arr, len(arr), lastParent)
	}
}

func siftDown(arr []int, heapSize, index int) {
	for index < heapSize {
		left := index*2 + 1
		right := index*2 + 2
		biggest := index
		if left < heapSize && arr[biggest] < arr[left] {
			biggest = left
		}
		if right < heapSize && arr[biggest] < arr[right] {
			biggest = right
		}
		if index == biggest {
			break
		}
		swap(arr, biggest, index)
		index = biggest
	}
}

func heapSort(arr []int) {
	// time = O(n)
	maxHeapify(arr)

	// time = O(n * logn)
	// space = O(1)
	for last := len(arr) - 1; last > 0; last-- {
		swap(arr, last, 0)
		siftDown(arr, last, 0)
	}
}
