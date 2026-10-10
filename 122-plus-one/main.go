package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{1, 9, 9})) // [2 0 0]
	fmt.Println(plusOne([]int{9, 9}))    // [1 0 0]
	fmt.Println(plusOne([]int{0}))       // [1]
}

// plusOne increments a nonempty decimal digit slice in place.
// Time: O(n). Extra space: O(1), excluding the O(n) result for all nines.
func plusOne(digits []int) []int {
	// Invariant: digits to the right are final zeros; a carry of 1 remains.
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// edge case: All digits were 9, so the result needs one more digit.
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
