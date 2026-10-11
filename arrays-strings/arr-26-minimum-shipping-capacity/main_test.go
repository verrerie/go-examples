package main

import "testing"

// Enumerate all cuts rather than using the implementation's greedy check.
func bruteCapacity(a []int, days int) int {
	if len(a) == 0 {
		return 0
	}
	best := int(^uint(0) >> 1)
	for mask := 0; mask < 1<<(len(a)-1); mask++ {
		used, load, largest := 1, 0, 0
		for i, v := range a {
			load += v
			if i == len(a)-1 || mask&(1<<i) != 0 {
				largest = max(largest, load)
				load = 0
				if i < len(a)-1 {
					used++
				}
			}
		}
		if used <= days {
			best = min(best, largest)
		}
	}
	return best
}
func TestCapacity(t *testing.T) {
	for n := 0; n <= 5; n++ {
		total := 1
		for range n {
			total *= 3
		}
		for code := 0; code < total; code++ {
			a := make([]int, n)
			x := code
			for i := range a {
				a[i] = x%3 + 1
				x /= 3
			}
			for days := 1; days <= n+1; days++ {
				if got, want := minCapacity(a, days), bruteCapacity(a, days); got != want {
					t.Fatal(a, days, got, want)
				}
			}
		}
	}
}
