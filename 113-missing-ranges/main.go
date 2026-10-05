package main

import "fmt"

func main() {
	nums := []int{0, 1, 3, 50, 75}
	fmt.Println(findMissingRanges(nums, 0, 99))
	fmt.Println(findMissingRanges([]int{-1}, -1, -1))
}

// time = O(n)
// space = O(1) excluding the result array
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	prev := lower - 1
	var current int
	result := [][]int{}
	recordGap := func(prev, current int) {
		result = append(result, []int{prev + 1, current - 1})
	}
	// time = O(n)
	for _, n := range nums {
		current = n
		if current-prev > 1 {
			recordGap(prev, current)
		}
		prev = current
	}
	current = upper + 1
	if current-prev > 1 {
		recordGap(prev, current)
	}
	return result
}
