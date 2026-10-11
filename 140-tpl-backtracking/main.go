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
func main() { fmt.Println(subsets([]int{1, 2, 3})) }
