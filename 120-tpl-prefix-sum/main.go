package main

import "fmt"

func main() {
	fmt.Println(prefixSum([]int{1, 2, 3, 4, 5, 6}))
}

func prefixSum(arr []int) []int {
	prefix := make([]int, len(arr))
	prefix[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	return prefix
}
