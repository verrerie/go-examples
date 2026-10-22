// Question: Given a nonempty integer slice guaranteed to contain a majority element,
// return that element. A majority element occurs more than len(nums)/2 times.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2, 1}))
	fmt.Println(majorityElement([]int{1, 0, 1, 0, 1, 0, 1}))
	fmt.Println(majorityElement([]int{1, 1, 0, 0, 2, 1, 1}))
}

func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
