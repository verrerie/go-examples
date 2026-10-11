package main

import "testing"

func TestCycle(t *testing.T) {
	for n := 0; n <= 8; n++ {
		nodes := make([]ListNode, n)
		for i := 0; i < n-1; i++ {
			nodes[i].Next = &nodes[i+1]
		}
		var head *ListNode
		if n > 0 {
			head = &nodes[0]
		}
		if hasCycle(head) {
			t.Fatal("acyclic", n)
		}
		if n > 0 {
			for i := range n {
				nodes[n-1].Next = &nodes[i]
				if !hasCycle(head) {
					t.Fatal("cycle", n, i)
				}
			}
		}
	}
}
