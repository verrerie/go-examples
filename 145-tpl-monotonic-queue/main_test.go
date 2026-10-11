package main

import (
	"slices"
	"testing"
)

func TestAgainstWindows(t *testing.T) {
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
			for k := 0; k <= n+1; k++ {
				want := []int{}
				if k > 0 && k <= n {
					for i := 0; i+k <= n; i++ {
						want = append(want, slices.Max(a[i:i+k]))
					}
				}
				if got := windowMax(a, k); !slices.Equal(got, want) {
					t.Fatal(a, k, got, want)
				}
			}
		}
	}
}
