package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time = O(h)
// space = O(1)
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var candidate *TreeNode
	var find func(*TreeNode)
	// time = O(h)
	find = func(node *TreeNode) {
		if node == nil {
			return
		} else if node.Val > p.Val {
			candidate = node
			find(node.Left)
		} else {
			find(node.Right)
		}
	}
	find(root)
	return candidate
}
