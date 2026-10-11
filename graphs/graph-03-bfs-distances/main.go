package main

import "fmt"

// distances returns unweighted shortest-path lengths; -1 means unreachable.
// Neighbor IDs must be valid. Invalid start leaves every distance at -1.
// Invariant: queued vertices have final distances in nondecreasing order.
// Time O(V+E); space O(V). Mark on enqueue, not dequeue.
func distances(graph [][]int, start int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = -1
	}
	if start < 0 || start >= len(graph) {
		return dist
	}
	dist[start] = 0
	queue := []int{start}
	for head := 0; head < len(queue); head++ {
		v := queue[head]
		for _, next := range graph[v] {
			if dist[next] == -1 {
				dist[next] = dist[v] + 1
				queue = append(queue, next)
			}
		}
	}
	return dist
}
func main() { fmt.Println(distances([][]int{{1, 2}, {3}, {3}, {}, {}}, 0)) }
