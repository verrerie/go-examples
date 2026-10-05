package main

import "fmt"

func main() {
	list := []int{2, 6, 3, 5, 7, 9, 0}
	quickSort(list)
	fmt.Println(list)
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	swap := func(l, r int) { nums[l], nums[r] = nums[r], nums[l] }
	// space(partition) = O(1)
	pivot := nums[0]
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
	// after movements, r is the last element of left side
	swap(0, r)
	// logn calls
	// stack space = O(logn) on average, O(n) in worst case
	quickSort(nums[:r])
	quickSort(nums[r+1:])
}
