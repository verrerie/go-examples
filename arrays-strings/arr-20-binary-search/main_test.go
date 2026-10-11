package main

import (
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	for _, a := range [][]int{nil, {1}, {1, 1, 1}, {-3, 0, 2, 2, 8}} {
		for target := -4; target <= 9; target++ {
			got := search(a, target)
			if slices.Contains(a, target) {
				if got < 0 || got >= len(a) || a[got] != target {
					t.Fatal(a, target, got)
				}
			} else if got != -1 {
				t.Fatal(a, target, got)
			}
		}
	}
}
