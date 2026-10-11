package main

import "fmt"

func main() {
	list := []int{1, 4, 2, 3, 4, 5, 6, 7, 6}
	fmt.Println(findPeakElement(list))
	list = []int{1, 2, 3}
	fmt.Println(findPeakElement(list))
	list = []int{1}
	fmt.Println(findPeakElement(list))
}

// time = O(logn)
// space = O(1)
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// every iteration we shrink a half
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// l is now equal to r
	return l
}
