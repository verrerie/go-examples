package main

import "fmt"

// time = O(n)
// space = O(1)
func findMaxConsecutiveOnes(nums []int) int {
	i, j := 0, 0
	maxL := 0
	for i < len(nums) && j < len(nums) {
		i = j
		for i < len(nums) && nums[i] != 1 {
			i++
		}
		j = i
		for j < len(nums) && nums[j] == 1 {
			j++
		}
		maxL = max(maxL, j-i)
	}
	return maxL
}
func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
