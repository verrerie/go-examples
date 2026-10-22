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
	interval := intervals[0]
	// time = O(n)
	for next := range intervals {
		if interval[1] >= intervals[next][0] {
			interval = []int{interval[0], max(interval[1], intervals[next][1])}
		} else {
			result = append(result, interval)
			interval = intervals[next]
		}
	}
	result = append(result, interval)
	// space = O(n)
	return result
}
