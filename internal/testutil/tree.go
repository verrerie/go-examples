package testutil

import "fmt"

// TreeNode is the binary tree node type used by the LeetCode tree exercises.
// Exercises can use type TreeNode = testutil.TreeNode to share this type.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MakeTree builds a binary tree from LeetCode-style level-order values.
// Each value must be an int or nil, where nil represents a missing child.
// Children are assigned left, then right, to each non-nil node in level order;
// missing nodes do not consume child entries. Trailing nil values may be omitted.
//
// For example, MakeTree(1, nil, 2, 3) creates:
//
//	1
//	 \
//	  2
//	 /
//	3
//
// Empty input or a nil root creates an empty tree. MakeTree panics if a value
// has an unsupported type or a non-nil entry has no parent in the tree.
func MakeTree(values ...any) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := makeTreeNode(values[0], 0)
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	index := 1
	for parent := 0; parent < len(queue) && index < len(values); parent++ {
		node := queue[parent]
		node.Left = makeTreeNode(values[index], index)
		index++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if index == len(values) {
			break
		}
		node.Right = makeTreeNode(values[index], index)
		index++
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	for ; index < len(values); index++ {
		if values[index] != nil {
			panic(fmt.Sprintf("testutil.MakeTree: entry %d has no parent", index))
		}
	}
	return root
}

func makeTreeNode(value any, index int) *TreeNode {
	if value == nil {
		return nil
	}
	n, ok := value.(int)
	if !ok {
		panic(fmt.Sprintf("testutil.MakeTree: entry %d must be an int or nil, got %T", index, value))
	}
	return &TreeNode{Val: n}
}

func ToSlice(t *TreeNode) []int {
	result := make([]int, 0)
	var visit func(*TreeNode)
	visit = func(tr *TreeNode) {
		if tr == nil {
			return
		}
		result = append(result, tr.Val)
		visit(tr.Left)
		visit(tr.Right)
	}
	visit(t)
	return result
}
