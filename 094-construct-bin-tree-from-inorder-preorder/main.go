package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type TreeNode = testutil.TreeNode

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Println(testutil.ToSlice(tree))
	fmt.Println(testutil.ToSlice(buildTree([]int{}, []int{})))
}

// time = O(n), space = O(n + h)
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	// O(n), space = O(n)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func([]int, int) *TreeNode
	// O(n), space = O(h)
	build = func(preorder []int, inStart int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		rootVal := preorder[0]
		rootIndexInorder := inorderMap[rootVal]

		leftSize := rootIndexInorder - inStart

		leftPreorder := preorder[1 : leftSize+1]
		rightPreorder := preorder[leftSize+1:]

		root := &TreeNode{Val: rootVal}
		root.Left = build(leftPreorder, inStart)
		root.Right = build(rightPreorder, rootIndexInorder+1)

		return root
	}

	return build(preorder, 0)
}
