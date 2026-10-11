package main

import (
	"slices"
	"testing"
)

func TestNextSmaller(t *testing.T) {
	for n := 0; n <= 7; n++ {
		total := 1
		for range n {
			total *= 3
		}
		for mask := 0; mask < total; mask++ {
			a, want := make([]int, n), make([]int, n)
			x := mask
			for i := range a {
				a[i] = x % 3
				x /= 3
			}
			for i := range a {
				want[i] = -1
				for j := i + 1; j < n; j++ {
					if a[j] < a[i] {
						want[i] = j
						break
					}
				}
			}
			if got := nextSmaller(a); !slices.Equal(got, want) {
				t.Fatalf("%v: got %v want %v", a, got, want)
			}
		}
	}
}
