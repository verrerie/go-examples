package main

import "testing"

func TestPiece(t *testing.T) {
	for n := 0; n <= 4; n++ {
		total := 1
		for range n {
			total *= 5
		}
		for code := 0; code < total; code++ {
			a := make([]int, n)
			x := code
			for i := range a {
				a[i] = x % 5
				x /= 5
			}
			for k := 1; k <= 10; k++ {
				want := 0
				for size := 1; size <= 4; size++ {
					count := 0
					for _, v := range a {
						count += v / size
					}
					if count >= k {
						want = size
					}
				}
				if got := maxPiece(a, k); got != want {
					t.Fatal(a, k, got, want)
				}
			}
		}
	}
}
