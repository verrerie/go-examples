// Question: Given a binary slice containing only 0 and 1, return the length of the
// longest contiguous run of 1s. Return 0 for an empty slice.

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
