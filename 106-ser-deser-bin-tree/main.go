package main

import (
	"fmt"
	"strconv"
	"strings"

	"go-examples/internal/testutil"
)

func main() {
	tree := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7, 8)
	codec := Constructor()
	serialized := codec.serialize(tree)
	fmt.Println(serialized)
	deserialized := codec.deserialize(serialized)
	reSerialized := codec.serialize(deserialized)
	fmt.Println(reSerialized)
	fmt.Println(reSerialized == serialized)
}

type TreeNode = testutil.TreeNode

type Codec struct{}

const SEP = ","

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// time = O(n) as each node is visited once
// space = O(n) as the string builder contains all the nodes
func (codec *Codec) serialize(root *TreeNode) string {
	build := strings.Builder{}
	var ser func(*TreeNode)
	ser = func(tree *TreeNode) {
		if tree == nil {
			build.WriteString("nil")
			return
		}
		// preorder ser
		build.WriteString(strconv.Itoa(tree.Val))
		build.WriteString(SEP)
		ser(tree.Left)
		build.WriteString(SEP)
		ser(tree.Right)
	}
	ser(root)
	return build.String()
}

// Deserializes your encoded data to tree.
// time = O(n), as every token is read once
// space = O(n+h) = O(n), n: tokens array is linear to the number of nodes, h: the stack depth, worst case h=n
func (codec *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	tokens := strings.Split(data, ",")
	var deser func(int) (*TreeNode, int)
	deser = func(start int) (*TreeNode, int) {
		if tokens[start] == "nil" {
			return nil, 1
		}
		val, err := strconv.Atoi(tokens[start])
		if err != nil {
			panic(fmt.Sprintln("invalid token encoutered:", tokens[start]))
		}
		// the ser string is in preorder
		root := &TreeNode{Val: val}
		left, numTokensLeft := deser(start + 1)
		root.Left = left
		right, numTokensRight := deser(start + 1 + numTokensLeft)
		root.Right = right
		return root, 1 + numTokensLeft + numTokensRight
	}
	result, _ := deser(0)
	return result
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
