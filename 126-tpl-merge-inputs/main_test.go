package main

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	for _, tc := range []struct{ a, b, want []int }{{nil, nil, []int{}}, {[]int{1}, nil, []int{1}}, {nil, []int{2}, []int{2}}, {[]int{-2, 1, 1}, []int{-1, 1, 3}, []int{-2, -1, 1, 1, 1, 3}}} {
		a, b := slices.Clone(tc.a), slices.Clone(tc.b)
		if got := mergeSorted(tc.a, tc.b); !slices.Equal(got, tc.want) {
			t.Fatalf("got %v want %v", got, tc.want)
		}
		if !slices.Equal(a, tc.a) || !slices.Equal(b, tc.b) {
			t.Fatal("mutated input")
		}
	}
}
