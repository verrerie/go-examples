package testutil

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMakeTree(t *testing.T) {
	cases := []struct {
		name   string
		values []any
		want   *TreeNode
	}{
		{name: "empty"},
		{name: "nil root", values: []any{nil}},
		{name: "empty with trailing nils", values: []any{nil, nil, nil}},
		{name: "zero is a node", values: []any{0}, want: &TreeNode{Val: 0}},
		{
			name:   "balanced",
			values: []any{1, 2, 3, 4, 5, 6, 7},
			want: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
			},
		},
		{
			name:   "compact sparse format",
			values: []any{1, nil, 2, 3},
			want:   &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
		},
		{
			name:   "interior gaps",
			values: []any{1, 2, 3, nil, 4, 5, nil, 6, nil, nil, 7},
			want: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}},
			},
		},
		{
			name:   "left chain",
			values: []any{1, 2, nil, 3, nil, 4},
			want:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}},
		},
		{
			name:   "right chain",
			values: []any{1, nil, 2, nil, 3, nil, 4},
			want:   &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}},
		},
		{
			name:   "negative and duplicate values",
			values: []any{-1, 0, -1},
			want:   &TreeNode{Val: -1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: -1}},
		},
		{
			name:   "trailing nils",
			values: []any{1, nil, nil, nil, nil},
			want:   &TreeNode{Val: 1},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MakeTree(tc.values...); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MakeTree(%v) produced an unexpected tree", tc.values)
			}
		})
	}
}

func TestMakeTreeRejectsInvalidInput(t *testing.T) {
	cases := []struct {
		name   string
		values []any
		want   string
	}{
		{"invalid root type", []any{"1"}, "entry 0 must be an int or nil"},
		{"invalid child type", []any{1, false}, "entry 1 must be an int or nil"},
		{"floating point", []any{1, 2.5}, "entry 1 must be an int or nil"},
		{"child of nil root", []any{nil, 1}, "entry 1 has no parent"},
		{"unreachable child", []any{1, nil, nil, 2}, "entry 3 has no parent"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				got := recover()
				if got == nil || !strings.Contains(fmt.Sprint(got), tc.want) {
					t.Errorf("panic = %v, want message containing %q", got, tc.want)
				}
			}()
			MakeTree(tc.values...)
		})
	}
}

func ExampleMakeTree() {
	root := MakeTree(1, nil, 2, 3)
	fmt.Println(root.Val, root.Right.Val, root.Right.Left.Val)
	// Output: 1 2 3
}
