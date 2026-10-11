package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestTopKAgainstSort(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for n := 0; n < 40; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = r.Intn(11) - 5
		}
		sorted := slices.Clone(a)
		slices.Sort(sorted)
		slices.Reverse(sorted)
		for k := -1; k <= n+1; k++ {
			want := sorted[:max(0, min(k, n))]
			if got := topK(a, k); !slices.Equal(got, want) {
				t.Fatalf("%v k=%d got %v want %v", a, k, got, want)
			}
		}
	}
}
