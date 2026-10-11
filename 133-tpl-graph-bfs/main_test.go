package main

import (
	"slices"
	"testing"
)

func TestDistances(t *testing.T) {
	g := [][]int{{0, 1, 2, 1}, {3}, {3}, {1}, {}}
	if got := distances(g, 0); !slices.Equal(got, []int{0, 1, 1, 2, -1}) {
		t.Fatal(got)
	}
	if got := distances(g, -1); !slices.Equal(got, []int{-1, -1, -1, -1, -1}) {
		t.Fatal(got)
	}
	if len(distances(nil, 0)) != 0 {
		t.Fatal("empty")
	}
}
