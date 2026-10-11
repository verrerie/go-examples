package main

import "testing"

func TestAgainstAllRanges(t *testing.T) {
	for n := 0; n <= 6; n++ {
		total := 1
		for range n {
			total *= 3
		}
		for mask := 0; mask < total; mask++ {
			a := make([]int, n)
			x := mask
			for i := range a {
				a[i] = x%3 - 1
				x /= 3
			}
			for k := -3; k <= 3; k++ {
				want := 0
				for i := range a {
					sum := 0
					for j := i; j < len(a); j++ {
						sum += a[j]
						if sum == k {
							want++
						}
					}
				}
				if got := countSum(a, k); got != want {
					t.Fatalf("%v k=%d got %d want %d", a, k, got, want)
				}
			}
		}
	}
}
