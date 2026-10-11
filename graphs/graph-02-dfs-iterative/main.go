package main

import "fmt"

// reachable counts vertices reachable from start. Neighbor IDs must be valid.
// Invariant: every stacked vertex is marked and has been pushed exactly once.
// Time O(V+E); space O(V). A stack gives depth-first exploration.
func reachable(graph [][]int, start int) int {
	if start < 0 || start >= len(graph) {
		return 0
	}
	seen := make([]bool, len(graph))
	seen[start] = true
	stack := []int{start}
	count := 0
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		for _, next := range graph[v] {
			if !seen[next] {
				seen[next] = true
				stack = append(stack, next)
			}
		}
	}
	return count
}
func main() { fmt.Println(reachable([][]int{{1, 2}, {2}, {0}, {}}, 0)) }
