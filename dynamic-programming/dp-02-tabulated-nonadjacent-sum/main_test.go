package main

import "testing"

func TestAgainstSubsets(t *testing.T) {
	for n := 0; n <= 7; n++ {
		total := 1
		for range n {
			total *= 3
		}
		for code := 0; code < total; code++ {
			a := make([]int, n)
			x := code
			for i := range a {
				a[i] = x%3 - 1
				x /= 3
			}
			want := 0
			for mask := 0; mask < 1<<n; mask++ {
				if mask&(mask<<1) != 0 {
					continue
				}
				sum := 0
				for i, v := range a {
					if mask&(1<<i) != 0 {
						sum += v
					}
				}
				want = max(want, sum)
			}
			if got := maxNonAdjacent(a); got != want {
				t.Fatal(a, got, want)
			}
		}
	}
}
