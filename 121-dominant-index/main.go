package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{1, 2, 4}))
	fmt.Println(dominantIndex([]int{}))
	fmt.Println(dominantIndex([]int{1}))
	fmt.Println(dominantIndex([]int{1, 2}))
	fmt.Println(dominantIndex([]int{0, 0}))
}

// time = O(n)
// space = O(1)
func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	curr := 0
	prev := 1
	if nums[1] >= nums[0] {
		curr = 1
		prev = 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[curr] {
			prev = curr
			curr = i
		} else if nums[i] > nums[prev] {
			prev = i
		}
	}
	if nums[curr] >= 2*nums[prev] {
		return curr
	}
	return -1
}
