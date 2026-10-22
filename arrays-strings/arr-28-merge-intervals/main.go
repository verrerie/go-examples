// Question: Given a nonempty slice of closed intervals [start, end], with start <= end,
// merge all overlapping intervals and return the resulting disjoint intervals
// in start order. Intervals sharing an endpoint overlap.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals))

	intervals = append(intervals, []int{17, 19}, []int{18, 22})
	fmt.Println(merge(intervals))
}

// time = O(n * logn)
// space = O(n)
func merge(intervals [][]int) [][]int {
	// time = O(n * logn)
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	result := [][]int{}
	// union is the current open union with earliest start
	union := intervals[0]
	// time = O(n)
	for next := range intervals {
		if union[1] >= intervals[next][0] {
			// case overlap
			union = []int{union[0], max(union[1], intervals[next][1])}
		} else {
			// record last union and start new union
			result = append(result, union)
			union = intervals[next]
		}
	}
	result = append(result, union)
	// space = O(n)
	return result
}
