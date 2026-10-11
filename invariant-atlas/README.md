# Invariant Atlas

A concise study website for exercises 082–145.

From this directory, run:

```sh
node build.mjs
node serve.mjs
```

Open http://localhost:4173. You can also open the generated `dist/index.html` directly. No package installation is required.

- Search by pattern, keyword, or exercise number.
- Select a card for the invariant, preservation step, stopping rule, example, cost, and source.
- Use Recall mode to hide answers. Review marks stay in the current browser.
- Prefix sum and three-way partition have interactive traces.

Edit `src/data.js` for study notes, `src/style.css` for design, and `src/app.js` for interactions. Run `node build.mjs` after changes. Git tracks `src/`; generated `dist/` files are ignored. Fonts have system fallbacks.

The build also copies Go source snapshots from the parent repository. To refresh only those snapshots, run `node sync-sources.mjs`. Neither command regenerates the editorial notes. Update the visible snapshot date in `src/index.html` when refreshing. The Go implementations are not modified.

The prose uses short, direct technical English inspired by STE. It is not a formally certified ASD-STE100 document.

## Cheatsheet coverage

All 25 code templates and both extensions from the [LeetCode cheatsheet](https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/) have invariant notes and Go implementations. Existing implementations are reused; folders 126–145 supply missing runnable examples with tests. Folder 119 still contains the original generic scaffolds. The two-input implementation is in 126.

Run the new examples and tests from the repository root:

```sh
go run ./144-tpl-dijkstra
go test ./12[6-9]-tpl-* ./13*-tpl-* ./14*-tpl-*
go vet ./12[6-9]-tpl-* ./13*-tpl-* ./14*-tpl-*
```

| Cheatsheet pattern | Atlas note | Go source |
| --- | --- | --- |
| Two pointers: opposite ends | Opposite pointers: sorted pair sum | [119](../119-tpl-two-pointers/main.go), [82](../082-3sum/main.go) |
| Two pointers: exhaust two inputs | Two input scans | [119](../119-tpl-two-pointers/main.go), [126](../126-tpl-merge-inputs/main.go) |
| Sliding window | Longest window: restore validity | [85](../085-longest-sub-str/main.go), [119](../119-tpl-two-pointers/main.go) |
| Prefix sum | Prefix sum | [120](../120-tpl-prefix-sum/main.go) |
| Efficient string building | Build a string efficiently | [88](../088-count-and-say/main.go), [106](../106-ser-deser-bin-tree/main.go) |
| Fast and slow linked-list pointers | Fast and slow: find the middle | [127](../127-tpl-fast-slow/main.go) |
| Reverse a linked list | Reverse a linked list | [128](../128-tpl-reverse-list/main.go) |
| Count subarrays with an exact criterion | Count ranges with prefix frequencies | [129](../129-tpl-prefix-counts/main.go) |
| Monotonic increasing stack | Monotonic increasing stack | [130](../130-tpl-monotonic-stack/main.go) |
| Tree DFS: recursive | Tree DFS: recursive | [92](../092-inorder-traversal/main.go), [106](../106-ser-deser-bin-tree/main.go) |
| Tree DFS: iterative | Inorder: explicit stack | [92](../092-inorder-traversal/main.go) |
| Tree BFS | BFS by level | [93](../093-zigzag-level/main.go), [95](../095-populate-next-right-ptr/main.go) |
| Graph DFS: recursive | Graph DFS: recursive | [131](../131-tpl-graph-dfs-recursive/main.go) |
| Graph DFS: iterative | Graph DFS: iterative | [132](../132-tpl-graph-dfs-iterative/main.go) |
| Graph BFS | Graph BFS: shortest unweighted paths | [133](../133-tpl-graph-bfs/main.go) |
| Top k with a heap | Keep top k with a heap | [134](../134-tpl-top-k-heap/main.go) |
| Binary search | Binary search: exact match | [135](../135-tpl-binary-search/main.go) |
| Left insertion point | Binary search: lower bound | [136](../136-tpl-lower-bound/main.go) |
| Right insertion point | Binary search: upper bound | [137](../137-tpl-upper-bound/main.go) |
| Search for a minimum feasible answer | Binary search on the minimum answer | [138](../138-tpl-search-minimum/main.go) |
| Search for a maximum feasible answer | Binary search on the maximum answer | [139](../139-tpl-search-maximum/main.go) |
| Backtracking | Backtracking: choose, explore, undo | [140](../140-tpl-backtracking/main.go) |
| Top-down dynamic programming | Dynamic programming: memoization | [141](../141-tpl-dp-memo/main.go) |
| Build a trie | Trie: share word prefixes | [143](../143-tpl-trie/main.go) |
| Dijkstra | Dijkstra: nonnegative shortest paths | [144](../144-tpl-dijkstra/main.go) |
| Extension: bottom-up dynamic programming | Dynamic programming: bottom-up table | [142](../142-tpl-dp-table/main.go) |
| Extension: monotonic queue | Monotonic queue: window maximum | [145](../145-tpl-monotonic-queue/main.go) |

## Pattern boundaries

- One sorted pair-sum card covers opposite pointers, Two Sum, and the inner 3Sum scan.
- Separate cards explain finding a list middle and detecting a cycle, partitioning and recursive quicksort, and inorder traversal and BST rank selection.
- Keep distinct contracts separate: longest versus shortest windows; occurrence matching versus insertion bounds; recursive versus iterative DFS; memoization versus tabulation; stack versus deque.
- Run scanning cards keep different output contracts: encoding all runs versus finding the longest run of ones.
- Prefix tables answer range queries. Prefix-frequency maps count matching ranges; they are separate patterns.
- Old two-pointer bookmarks redirect. Existing review marks migrate once when the site loads; newly split parts inherit the old mark.

Validate catalog links and review migration with `node --test tests/catalog.test.mjs`.
