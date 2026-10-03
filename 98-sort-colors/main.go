package main

import "fmt"

func main() {
	list := []int{1, 0, 2, 1, 0, 0, 2, 2}
	sortColors(list)
	fmt.Println(list)
}

// time = O(n), every iteration shrink the unprocessed space by 1
// space = O(1)
func sortColors(nums []int) {
	zeroIndex := 0
	twosIndex := len(nums) - 1
	swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
	n := 0
	for n <= twosIndex {
		switch nums[n] {
		case 0:
			swap(zeroIndex, n)
			zeroIndex++
			n++
		case 1:
			n++
		case 2:
			swap(twosIndex, n)
			twosIndex--
		}
	}
}
