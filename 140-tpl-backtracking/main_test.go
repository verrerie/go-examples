package main

import "testing"

func TestSubsets(t *testing.T) {
	for n := 0; n <= 8; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = i
		}
		out := subsets(a)
		if len(out) != 1<<n {
			t.Fatal("count", n)
		}
		seen := map[int]bool{}
		for _, s := range out {
			mask, prev := 0, -1
			for _, v := range s {
				if v <= prev || v >= n {
					t.Fatal("order", s)
				}
				mask |= 1 << v
				prev = v
			}
			if seen[mask] {
				t.Fatal("duplicate", s)
			}
			seen[mask] = true
		}
	}
	out := subsets([]int{1, 2})
	out[1][0] = 99
	if out[2][0] != 1 {
		t.Fatal("results alias")
	}
}
