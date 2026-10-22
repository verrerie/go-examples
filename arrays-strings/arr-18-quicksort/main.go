// Question: Given an integer slice, sort its values in nondecreasing order in place
// using quicksort. Empty and single-element slices are already sorted.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	list := []int{2, 6, 3, 5, 7, 9, 0}
	quickSort(list)
	fmt.Println(list)
	list = []int{1, 2, 3}
	quickSort(list)
	fmt.Println(list)
	list = []int{1}
	quickSort(list)
	fmt.Println(list)
	list = []int{}
	quickSort(list)
	fmt.Println(list)
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	swap := func(l, r int) { nums[l], nums[r] = nums[r], nums[l] }
	// space(partition) = O(1)
	pivotIndex := rand.IntN(len(nums))
	pivot := nums[pivotIndex]
	swap(0, pivotIndex)
	l, r := 1, len(nums)-1
	// time(partition) = O(n)
	for l <= r { // we still need to check nums[l] when l == r
		//[2, 0, 5^$, 7, 9, 3, 6]
		if nums[l] > pivot {
			swap(l, r)
			r--
		} else {
			l++
		}
	}
	// r is now the last element of left half
	swap(0, r)
	// logn calls
	// stack space = O(logn) on average, O(n) in worst case
	quickSort(nums[:r])
	quickSort(nums[r+1:])
}
