// Question: Given a nonempty integer slice, return a slice of the same length where
// result[i] is the sum of input elements from index 0 through i, inclusive.

package main

import "fmt"

// prefix sum pattern
func prefixSum(arr []int) []int {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	return prefix
}
func main() {
	fmt.Println(prefixSum([]int{1, 2, 3, 4, 5, 6}))
}
