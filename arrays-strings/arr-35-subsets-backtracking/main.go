package main

import "fmt"

// subsets enumerates index-based subsets; use distinct values for unique value sets.
// Invariant: path contains exactly the choices on the current recursion path.
// Undo each choice before exploring its sibling. Copy path when saving a result.
// Time and output O(n*2^n); extra path and recursion O(n).
func subsets(nums []int) [][]int {
	out := [][]int{}
	path := []int{}
	var visit func(int)
	visit = func(start int) {
		saved := append([]int{}, path...)
		out = append(out, saved)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			visit(i + 1)
			path = path[:len(path)-1]
		}
	}
	visit(0)
	return out
}

func subsets2(nums []int) [][]int {
	out := [][]int{}
	var visit func(int, []int)
	visit = func(start int, path []int) {
		saved := append([]int{}, path...)
		out = append(out, saved)
		for i := start; i < len(nums); i++ {
			visit(i+1, append(path, nums[i]))
		}
	}
	visit(0, make([]int, 0, len(nums)))
	return out
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets2([]int{4, 5, 6, 7, 8, 9}))
}
