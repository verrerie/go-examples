package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type TreeNode = testutil.TreeNode

func main() {
	balanced := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(kthSmallest(balanced, 3))
}

// O(h + k), space = O(h)
func kthSmallest(root *TreeNode, k int) int {
	// space = O(h), h=n in worst case
	stack := make([]*TreeNode, 0)
	// O(h) and h = n in worst scenario
	stack = pushLeftPath(root, stack)
	// we should exit earlier before reaching all nodes
	var current *TreeNode
	// O(k) as we pop k times, each node is pushed and popped at most once
	for len(stack) > 0 {
		current, stack = pop(stack)
		k--
		if k == 0 {
			return current.Val
		}
		stack = pushLeftPath(current.Right, stack)
	}
	panic("this shouldn't happen")
}

func pushLeftPath(node *TreeNode, stack []*TreeNode) []*TreeNode {
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}
	return stack
}

func pop(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	top := len(stack) - 1
	return stack[top], stack[:top]
}
