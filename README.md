# Go examples and algorithm patterns

A grouped library of 153 runnable Go examples. Language features follow [Go by Example](https://gobyexample.com); algorithm notes are in [Invariant Atlas](invariant-atlas/README.md).

## Find code

Each group has its own folder. Every implementation uses **prefix-number-name**, with numbering starting at **01** within that group. Two-digit numbers keep file-browser sorting correct.

| Group | Prefix | Examples | Index |
| --- | --- | ---: | --- |
| Go language and features | golang | 85 | [golang](golang/README.md) |
| Arrays and strings | arr | 35 | [arrays-strings](arrays-strings/README.md) |
| Linked lists | list | 6 | [linked-lists](linked-lists/README.md) |
| Trees and tries | tree | 8 | [trees](trees/README.md) |
| Graphs and grids | graph | 6 | [graphs](graphs/README.md) |
| Hash maps and counting | hash | 4 | [hash-maps](hash-maps/README.md) |
| Stacks and queues | sq | 3 | [stacks-queues](stacks-queues/README.md) |
| Heaps and priority queues | heap | 3 | [heaps](heaps/README.md) |
| Dynamic programming | dp | 2 | [dynamic-programming](dynamic-programming/README.md) |
| Bit operations | bit | 1 | [bit-operations](bit-operations/README.md) |

## Run and check

Run from the repository root:

~~~sh
go run ./golang/golang-01-hello-world
go run ./arrays-strings/arr-04-prefix-sum
go run ./graphs/graph-06-dijkstra
go test ./...
go vet ./...
node scripts/check-library.mjs
node invariant-atlas/build.mjs
node --test invariant-atlas/tests/catalog.test.mjs
~~~

File-I/O examples use paths relative to the working directory; run them from their own folder when using their sample files. Embedded assets, tests, and data files move with each example.

## Grouping rules

Choose the main data structure or technique used by the implementation. Keep algorithm names in references so searches for binary search, intervals, backtracking, or BFS still work. Use the group index for borderline cases; do not duplicate an implementation across groups.

Go language features have the golang prefix, including the five function topics that previously shared one file. The original unfinished two-pointer/window scaffolds remain beside their matching complete implementations and are marked as scaffolds.

The [code catalog](code-library.json) is the source of truth for paths and old-number aliases. The site shows full references and can search old exercise numbers. Existing invariant bookmarks and review marks are preserved.

To add an example, use the next number in its group and add one catalog entry. Run the library audit, rebuild the atlas, and update that group's index.
