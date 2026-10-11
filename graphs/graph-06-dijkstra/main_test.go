package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestShortestPaths(t *testing.T) {
	g := [][]Edge{{{1, 10}, {2, 1}}, {{3, 1}}, {{1, 1}, {3, 20}}, {}, {}}
	if got := shortestPaths(g, 0); !slices.Equal(got, []int64{0, 2, 1, 3, infinity}) {
		t.Fatal(got)
	}
	if len(shortestPaths(nil, 0)) != 0 {
		t.Fatal("empty")
	}
	g = [][]Edge{{{1, infinity - 2}}, {{2, 10}}, {}}
	if got := shortestPaths(g, 0); got[2] != infinity {
		t.Fatal("overflow", got)
	}
}
func TestAgainstFloydWarshall(t *testing.T) {
	r := rand.New(rand.NewSource(17))
	for trial := 0; trial < 100; trial++ {
		n := 1 + r.Intn(8)
		g := make([][]Edge, n)
		d := make([][]int64, n)
		for i := range n {
			d[i] = make([]int64, n)
			for j := range n {
				d[i][j] = infinity
			}
			d[i][i] = 0
		}
		for i := range n {
			for j := range n {
				if r.Intn(3) == 0 {
					w := int64(r.Intn(8))
					g[i] = append(g[i], Edge{j, w})
					d[i][j] = min(d[i][j], w)
				}
			}
		}
		for k := range n {
			for i := range n {
				for j := range n {
					if d[i][k] != infinity && d[k][j] != infinity {
						d[i][j] = min(d[i][j], d[i][k]+d[k][j])
					}
				}
			}
		}
		for source := range n {
			if got := shortestPaths(g, source); !slices.Equal(got, d[source]) {
				t.Fatal(g, source, got, d[source])
			}
		}
	}
}
func TestRejectNegative(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("negative weight accepted")
		}
	}()
	shortestPaths([][]Edge{{{0, -1}}}, 0)
}
