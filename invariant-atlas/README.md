# Invariant Atlas

69 invariant notes and 152 Go implementations in 10 groups. Dark GitHub styling; concise technical English.

From this directory:

~~~sh
node build.mjs
node serve.mjs
node --test tests/catalog.test.mjs
~~~

Open http://localhost:4173. Search by algorithm, group, prefix-number-name reference, or old exercise number. Group filters apply to both notes and the code index. Go language examples are available in the code index; they do not all have invariant notes.

Source paths come from [code-library.json](../code-library.json). Builds copy main.go and any teaching scaffold into a labeled text snapshot. Edit src/; dist/ is generated and ignored. Source snapshots use full references, for example sources/arr-04-prefix-sum.txt. Stable pattern IDs preserve bookmarks and browser review marks.

## Cheatsheet coverage

All 25 templates and both extensions from the [LeetCode cheatsheet](https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/) remain covered.

| Pattern | Reference |
| --- | --- |
| Two pointers: opposite ends | [arr-01-two-sum-sorted](../arrays-strings/arr-01-two-sum-sorted/main.go), [arr-16-three-sum](../arrays-strings/arr-16-three-sum/main.go) |
| Two pointers: exhaust two inputs | [arr-02-merge-sorted-inputs](../arrays-strings/arr-02-merge-sorted-inputs/main.go) |
| Sliding window | [arr-07-longest-unique-substring](../arrays-strings/arr-07-longest-unique-substring/main.go) |
| Prefix sum | [arr-04-prefix-sum](../arrays-strings/arr-04-prefix-sum/main.go) |
| Efficient string building | [arr-11-count-and-say](../arrays-strings/arr-11-count-and-say/main.go), [tree-06-serialize-deserialize](../trees/tree-06-serialize-deserialize/main.go) |
| Fast and slow linked-list pointers | [list-04-middle](../linked-lists/list-04-middle/main.go) |
| Reverse a linked list | [list-06-reverse-list](../linked-lists/list-06-reverse-list/main.go) |
| Count subarrays with an exact criterion | [hash-03-subarray-sum-count](../hash-maps/hash-03-subarray-sum-count/main.go) |
| Monotonic increasing stack | [sq-02-next-smaller-stack](../stacks-queues/sq-02-next-smaller-stack/main.go) |
| Tree DFS: recursive | [tree-01-inorder-traversal](../trees/tree-01-inorder-traversal/main.go), [tree-06-serialize-deserialize](../trees/tree-06-serialize-deserialize/main.go) |
| Tree DFS: iterative | [tree-01-inorder-traversal](../trees/tree-01-inorder-traversal/main.go) |
| Tree BFS | [tree-02-zigzag-level-order](../trees/tree-02-zigzag-level-order/main.go), [tree-04-connect-next-pointers](../trees/tree-04-connect-next-pointers/main.go) |
| Graph DFS: recursive | [graph-01-dfs-recursive](../graphs/graph-01-dfs-recursive/main.go) |
| Graph DFS: iterative | [graph-02-dfs-iterative](../graphs/graph-02-dfs-iterative/main.go) |
| Graph BFS | [graph-03-bfs-distances](../graphs/graph-03-bfs-distances/main.go) |
| Top k with a heap | [heap-01-top-k-elements](../heaps/heap-01-top-k-elements/main.go) |
| Binary search | [arr-20-binary-search](../arrays-strings/arr-20-binary-search/main.go) |
| Left insertion point | [arr-21-lower-bound](../arrays-strings/arr-21-lower-bound/main.go) |
| Right insertion point | [arr-22-upper-bound](../arrays-strings/arr-22-upper-bound/main.go) |
| Search for a minimum feasible answer | [arr-26-minimum-shipping-capacity](../arrays-strings/arr-26-minimum-shipping-capacity/main.go) |
| Search for a maximum feasible answer | [arr-27-maximum-piece-length](../arrays-strings/arr-27-maximum-piece-length/main.go) |
| Backtracking | [arr-35-subsets-backtracking](../arrays-strings/arr-35-subsets-backtracking/main.go) |
| Top-down dynamic programming | [dp-01-memoized-nonadjacent-sum](../dynamic-programming/dp-01-memoized-nonadjacent-sum/main.go) |
| Build a trie | [tree-08-trie](../trees/tree-08-trie/main.go) |
| Dijkstra | [graph-06-dijkstra](../graphs/graph-06-dijkstra/main.go) |
| Extension: bottom-up dynamic programming | [dp-02-tabulated-nonadjacent-sum](../dynamic-programming/dp-02-tabulated-nonadjacent-sum/main.go) |
| Extension: monotonic queue | [sq-03-window-maximum-deque](../stacks-queues/sq-03-window-maximum-deque/main.go) |

Distinct stopping rules remain separate: middle/cycle, partition/quicksort, inorder/rank, longest/shortest windows, exact-match/insertion searches, recursive/iterative DFS, and memoized/tabulated DP.

## Interview quick reference

The `#quick-reference` section summarizes [LeetCode resource 4725](https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4725/): searchable pattern cues, Go operation costs, input-size hints, and interview steps. Dynamic programming and backtracking are excluded from this summary. Existing catalog entries remain available. Cue links reuse stable invariant IDs and their Go source references.
