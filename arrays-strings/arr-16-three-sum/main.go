// Question: Given an integer slice, return all unique triplets of values whose sum is
// zero. Each triplet must use three distinct indices; duplicate values are
// allowed when present at separate indices. Result order does not matter.

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// sorted: -4, -1 -1, 0, 1, 2
	fmt.Println(threeSum([]int{-1, -1, -1}))
	fmt.Println(threeSum([]int{-1, -1, -1, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	// time = O(NlogN)
	slices.Sort(nums)
	result := [][]int{}

	// O(N)
	// two last nubmers cannot form an answer
	for i := 0; i < len(nums)-2; i++ {
		// when arr is sorted, we cannot get 0 with only positive nums
		if nums[i] > 0 {
			break
		}
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		// O(N)
		// no need to visit l == r because 1 number is not sufficient
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
				// skip duplicates
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				// skip duplicates
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}

		}
	}

	return result
}
