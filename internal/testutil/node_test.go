package testutil

import (
	"reflect"
	"testing"
)

func TestMakeNodeTree(t *testing.T) {
	cases := []struct {
		name   string
		values []any
		want   *Node
	}{
		{name: "empty"},
		{name: "nil root", values: []any{nil}},
		{name: "single", values: []any{0}, want: &Node{Val: 0}},
		{name: "sparse", values: []any{1, nil, 2, 3}, want: &Node{Val: 1, Right: &Node{Val: 2, Left: &Node{Val: 3}}}},
		{name: "negative and duplicate", values: []any{-1, 0, -1}, want: &Node{Val: -1, Left: &Node{Val: 0}, Right: &Node{Val: -1}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MakeNodeTree(tc.values...); !reflect.DeepEqual(got, tc.want) {
				t.Fatal("unexpected structure, values, or initial Next pointers")
			}
		})
	}
	a, b := MakeNodeTree(1, 2, 3), MakeNodeTree(1, 2, 3)
	if a == b || a.Left == b.Left || a.Right == b.Right {
		t.Fatal("separate fixtures share nodes")
	}
}
