package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int64
}
type entry struct {
	node     int
	distance int64
}
type distanceHeap []entry

func (h distanceHeap) Len() int           { return len(h) }
func (h distanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h distanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *distanceHeap) Push(v any)        { *h = append(*h, v.(entry)) }
func (h *distanceHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

const infinity int64 = math.MaxInt64

// shortestPaths requires valid vertex IDs, nonnegative weights, and shortest
// distances below infinity. infinity marks unreachable or unrepresentable distances.
// Invariant: popping a current minimum distance makes that distance final.
// Old heap entries are skipped; only a strict improvement adds a new entry.
// Lazy-heap time O(V + E log(E+1)); space O(V+E).
func shortestPaths(graph [][]Edge, source int) []int64 {
	for _, edges := range graph {
		for _, e := range edges {
			if e.To < 0 || e.To >= len(graph) || e.Weight < 0 {
				panic("invalid edge")
			}
		}
	}
	dist := make([]int64, len(graph))
	for i := range dist {
		dist[i] = infinity
	}
	if source < 0 || source >= len(graph) {
		return dist
	}
	dist[source] = 0
	h := &distanceHeap{{source, 0}}
	for h.Len() > 0 {
		current := heap.Pop(h).(entry)
		if current.distance != dist[current.node] {
			continue
		}
		for _, e := range graph[current.node] {
			if e.Weight >= infinity-current.distance {
				continue
			} // No signed overflow.
			candidate := current.distance + e.Weight
			if candidate < dist[e.To] {
				dist[e.To] = candidate
				heap.Push(h, entry{e.To, candidate})
			}
		}
	}
	return dist
}
func main() { fmt.Println(shortestPaths([][]Edge{{{1, 5}, {2, 1}}, {{3, 1}}, {{1, 1}}, {}}, 0)) }
