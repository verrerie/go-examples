package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// sorted: -4, -1 -1, 0, 1, 2
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	// O(N)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		// O(N)
		for l < r {
			total := nums[i] + nums[l] + nums[r]
			switch {
			case total < 0:
				l++
			case total > 0:
				r--
			case total == 0:
				// found
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}

		}
	}

	return result
}
