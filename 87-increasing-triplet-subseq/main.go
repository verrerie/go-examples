package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{2, 1, 6, 3, 9, 5}))
	fmt.Println(increasingTriplet([]int{2, 2, 2}))
	fmt.Println(increasingTriplet([]int{2, 1, 2}))
	fmt.Println(increasingTriplet([]int{2}))
	fmt.Println(increasingTriplet([]int{2, 2}))
	fmt.Println(increasingTriplet([]int{}))
}

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, x := range nums {
		if x <= first {
			first = x
		} else if x <= second {
			second = x
		} else {
			return true
		}
	}
	return false
}
