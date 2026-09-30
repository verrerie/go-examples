package main

import (
	"fmt"

	"go-examples/internal/testutil"
)

type TreeNode = testutil.TreeNode

func main() {
	balanced := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(balanced)
}
