package main

import "fmt"

func main() {
	list := []int{4, 5, 6, 7, 0, 1, 2, 3}
	fmt.Println(search(list, 1))
	fmt.Println(search(list, 2))
	fmt.Println(search(list, 3))
	fmt.Println(search(list, 4))
	fmt.Println(search(list, 5))
	fmt.Println(search(list, 6))
	fmt.Println(search(list, 7))
	fmt.Println(search(list, 0))
}

// time = O(logn)
// space = O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else {
			// time = O(logn), every iteration the size shrinks a half
			if nums[left] <= nums[mid] {
				// left is sorted
				if nums[left] <= target && target < nums[mid] {
					// target is on the left
					right = mid - 1
				} else {
					// target is on the right
					left = mid + 1
				}
			} else {
				// right is sorted
				if nums[mid] < target && target <= nums[right] {
					// target is on the right
					left = mid + 1
				} else {
					// target is on the left
					right = mid - 1
				}
			}
		}
	}
	return -1
}
