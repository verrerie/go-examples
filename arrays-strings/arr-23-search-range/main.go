package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 5, 5, 6, 7}
	fmt.Println(searchRange(list, 5))
	list = []int{}
	fmt.Println(searchRange(list, 1))
	list = []int{1}
	fmt.Println(searchRange(list, 1))
}

// time = O(logn)
// space = O(1)
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	start := -1
	// shrink half each iteration, time = O(logn)
	// space = O(1)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// we have a match
			start = mid
			right = mid - 1
		}
	}
	if start == -1 {
		return []int{-1, -1}
	}
	end := -1
	left, right = start, len(nums)-1
	// similar as start search
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// we have a match
			end = mid
			left = mid + 1
		}

	}

	return []int{start, end}
}
