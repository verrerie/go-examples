package main

import (
	"go-examples/internal/testutil"
	"testing"
)

func TestConnect(t *testing.T) {
	large := make([]any, 4095)
	for i := range large {
		large[i] = i%2001 - 1000
	}
	cases := []struct {
		name   string
		values []any
	}{
		{name: "empty"},
		{name: "single", values: []any{42}},
		{name: "two levels", values: []any{1, 2, 3}},
		{name: "cross parent links", values: []any{1, 2, 3, 4, 5, 6, 7}},
		{name: "four levels", values: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
		{name: "duplicate values", values: []any{0, -1, -1, 0, 0, 0, 0}},
		{name: "maximum size", values: large},
	}
	for _, implementation := range []struct {
		name string
		fn   func(*Node) *Node
	}{{"queue", connect}, {"constant space", connectV2}} {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					root := testutil.MakeNodeTree(tc.values...)
					// Capture levels and original contents before the implementation runs.
					var levels [][]*Node
					original := make(map[*Node]Node)
					if root != nil {
						levels = append(levels, []*Node{root})
					}
					for depth := 0; depth < len(levels); depth++ {
						var next []*Node
						for _, node := range levels[depth] {
							original[node] = *node
							if node.Left != nil {
								next = append(next, node.Left)
							}
							if node.Right != nil {
								next = append(next, node.Right)
							}
						}
						if len(next) > 0 {
							levels = append(levels, next)
						}
					}
					if got := implementation.fn(root); got != root {
						t.Fatal("returned a different root")
					}
					for depth, level := range levels {
						for i, node := range level {
							var want *Node
							if i+1 < len(level) {
								want = level[i+1]
							}
							if node.Next != want {
								t.Fatalf("level %d node %d: Next = %p, want %p", depth, i, node.Next, want)
							}
							before := original[node]
							if node.Val != before.Val || node.Left != before.Left || node.Right != before.Right {
								t.Fatal("modified original tree contents")
							}
						}
					}
				})
			}
		})
	}
}
