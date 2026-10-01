package testutil

// Node is a binary tree node with a pointer to its next neighbor on the same level.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// MakeNodeTree builds a tree using the same level-order format as MakeTree.
// All Next pointers start as nil. Each call creates fresh nodes.
func MakeNodeTree(values ...any) *Node {
	var convert func(*TreeNode) *Node
	convert = func(tree *TreeNode) *Node {
		if tree == nil {
			return nil
		}
		return &Node{Val: tree.Val, Left: convert(tree.Left), Right: convert(tree.Right)}
	}
	return convert(MakeTree(values...))
}
