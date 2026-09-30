package main

import "go-examples/internal/testutil"

type TreeNode = testutil.TreeNode

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 100)

	var visit func(*TreeNode)

	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			visit(node.Left)
		}
		result = append(result, node.Val)
		if node.Right != nil {
			visit(node.Right)
		}
	}

	visit(root)

	return result
}
