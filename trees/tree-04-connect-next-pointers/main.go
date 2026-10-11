package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type Node = testutil.Node

func main() {
	root := testutil.MakeNodeTree(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(connectV2(root))
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 1)
	queue[0] = root
	for len(queue) > 0 {
		levelSize := len(queue)
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		for i := range levelSize {
			if i+1 < levelSize {
				queue[i].Next = queue[i+1]
			} else {
				queue[i].Next = nil
			}
		}
		queue = queue[levelSize:]

	}
	return root
}

func connectV2(root *Node) *Node {
	if root == nil {
		return root
	}
	current := root
	// iterate when children exist
	for current.Left != nil {
		// current level's Next are connected
		head := current
		for head != nil {
			// if has children
			if head.Left != nil {
				head.Left.Next = head.Right
			}
			prev := head
			head = head.Next
			if head != nil {
				prev.Right.Next = head.Left
			}
		}
		// level connected
		// move to next level
		current = current.Left
	}
	return root
}
