package main

import "testing"

func TestReachable(t *testing.T) {
	g := [][]int{{0, 1, 1, 2}, {2}, {0}, {4}, {}}
	for _, tc := range []struct{ start, want int }{{0, 3}, {3, 2}, {4, 1}, {-1, 0}, {5, 0}} {
		if got := reachable(g, tc.start); got != tc.want {
			t.Fatal(tc, got)
		}
	}
	if reachable(nil, 0) != 0 {
		t.Fatal("empty")
	}
}
