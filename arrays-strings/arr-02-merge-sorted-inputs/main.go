package main

import "fmt"

// mergeSorted merges ascending inputs without changing them.
// Invariant: out is the sorted merge of a[:i] and b[:j].
// Time O(m+n); output O(m+n), extra state O(1).
func mergeSorted(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}
func main() { fmt.Println(mergeSorted([]int{1, 3, 5}, []int{2, 3, 6})) }
