package main

import "testing"

func TestPointers(t *testing.T) {
	for n := 0; n <= 8; n++ {
		nodes := make([]ListNode, n)
		for i := 0; i < n-1; i++ {
			nodes[i].Next = &nodes[i+1]
		}
		var head, want *ListNode
		if n > 0 {
			head = &nodes[0]
			want = &nodes[n/2]
		}
		if middle(head) != want || hasCycle(head) {
			t.Fatal("acyclic", n)
		}
		if n > 0 {
			for entry := range n {
				nodes[n-1].Next = &nodes[entry]
				if !hasCycle(head) {
					t.Fatal("missed cycle", n, entry)
				}
			}
		}
	}
}
