package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type TreeNode = testutil.TreeNode

func main() {
	balanced := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(inorderTraversal(balanced))
	fmt.Println(inorderTraversalV2(balanced))
}

func inorderTraversalV2(root *TreeNode) []int {
	result := make([]int, 0, 100)

	var visit func(*TreeNode)

	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Left)
		result = append(result, node.Val)
		visit(node.Right)
	}

	visit(root)

	return result
}

// O(N * LogN): the number of nodes, as we visit each node once.
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	// space = O(N), the worst case is when the tree is a linked list
	stack := make([]*TreeNode, 0)
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		stack, current = pop(stack)
		result = append(result, current.Val)
		current = current.Right

	}
	return result
}

func pop(stack []*TreeNode) ([]*TreeNode, *TreeNode) {
	last := len(stack) - 1
	node := stack[last]
	stack = stack[:last]
	return stack, node
}
