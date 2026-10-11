package main

import "fmt"

func pivotIndex(nums []int) int {
	// space = O(1)
	if len(nums) == 0 {
		return -1
	}
	sum := 0
	// time = O(n)
	for _, n := range nums {
		sum = sum + n
	}

	left := 0
	right := sum
	// time = O(n)
	for pivot := range nums {
		if pivot > 0 {
			left = left + nums[pivot-1]
		}
		right = right - nums[pivot]
		fmt.Println("pivot,left,right=", pivot, left, right)
		if left == right {
			return pivot
		}
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
