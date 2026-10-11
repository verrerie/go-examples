package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{1, 2, 4}))
	fmt.Println(dominantIndex([]int{}))
	fmt.Println(dominantIndex([]int{1}))
	fmt.Println(dominantIndex([]int{1, 2}))
	fmt.Println(dominantIndex([]int{1, 0}))
}

// time = O(n)
// space = O(1)
func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	maxIndex := 0
	secondMaxIndex := 1
	if nums[1] >= nums[0] {
		maxIndex = 1
		secondMaxIndex = 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			secondMaxIndex = maxIndex
			maxIndex = i
		} else if nums[i] > nums[secondMaxIndex] {
			secondMaxIndex = i
		}
	}
	if nums[maxIndex] >= 2*nums[secondMaxIndex] {
		return maxIndex
	}
	return -1
}
