# Shared exercise helpers

Import `go-examples/internal/testutil` from any exercise in this repository.

## Build a binary tree

Use the shared node type through an alias so your function signatures keep the
same `*TreeNode` spelling as LeetCode:

```go
import "go-examples/internal/testutil"

type TreeNode = testutil.TreeNode
```

Inside `main` or a test, build a tree with values in LeetCode's level order:

```go
root := testutil.MakeTree(1, nil, 2, 3)
//     1
//      \
//       2
//      /
//     3

balanced := testutil.MakeTree(1, 2, 3, 4, 5, 6, 7)
empty := testutil.MakeTree()
```

Each argument is an `int` or `nil`. Starting at the root, supply the left and
right children of each existing node in level order. Missing nodes do not get
child entries of their own; this is why `3` in the first example becomes `2`'s
left child. You can omit trailing `nil` entries. Zero, negative values, and
duplicate values are all ordinary node values.

Malformed fixtures, such as a string value or a child after an empty root, panic
with the entry index. Each call creates fresh nodes, so exercises can modify a
tree without affecting trees created by earlier calls.

## Capture example output

`CaptureStdout(t, fn)` captures everything printed while `fn` runs.
`CaptureLines(t, fn)` returns the trimmed, non-empty output lines.
