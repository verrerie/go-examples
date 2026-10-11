package main

import "fmt"

// reachable counts vertices reachable from start, not disconnected components.
// Input: adjacency lists with valid vertex IDs. Invalid start returns 0.
// Invariant: seen vertices have already been discovered; each is visited once.
// Time O(V+E); space O(V), including recursion.
func reachable(graph [][]int, start int) int {
	if start < 0 || start >= len(graph) {
		return 0
	}
	seen := make([]bool, len(graph))
	var visit func(int) int
	visit = func(v int) int {
		seen[v] = true
		count := 1
		for _, next := range graph[v] {
			if !seen[next] {
				count += visit(next)
			}
		}
		return count
	}
	return visit(start)
}
func main() { fmt.Println(reachable([][]int{{1, 2}, {2}, {0}, {}}, 0)) }
