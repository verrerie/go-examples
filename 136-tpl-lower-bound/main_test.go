package main

import "testing"

func TestLowerBound(t *testing.T) {
	for _, a := range [][]int{nil, {1}, {1, 1, 1}, {-3, 0, 2, 2, 8}} {
		for target := -4; target <= 9; target++ {
			want := 0
			for want < len(a) && a[want] < target {
				want++
			}
			if got := lowerBound(a, target); got != want {
				t.Fatal(a, target, got, want)
			}
		}
	}
}
