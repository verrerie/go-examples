package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type TreeNode = testutil.TreeNode

func main() {
	balanced := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(zigzagLevelOrder(balanced))
	fmt.Println(zigzagLevelOrder(nil))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0, 2000)
	if root == nil {
		return result
	}
	leftToRight := true
	queue := []*TreeNode{root}

	// stop if queue is empty, meaning no child left to visit
	for len(queue) > 0 {
		// we need to visit exactly levelSize nodes at current level
		levelSize := len(queue)
		level := make([]int, levelSize)
		for i := range levelSize {
			index := i
			// in case of directly flip, we read result backwards, without chaning the order of the queue
			if !leftToRight {
				index = levelSize - i - 1
			}
			// visit logic starts here
			visiting := queue[i]
			// put the Val in the right slot
			level[index] = visiting.Val
			// enqueue non nil children
			if visiting.Left != nil {
				queue = append(queue, visiting.Left)
			}
			if visiting.Right != nil {
				queue = append(queue, visiting.Right)
			}
			// visit logic done
		}
		// level complete, append to result
		result = append(result, level)
		// remove parent prefix before next level
		queue = queue[levelSize:]
		// flip the direction
		leftToRight = !leftToRight
	}

	return result
}
