package main

import "testing"

func TestReverse(t *testing.T) {
	for n := 0; n <= 8; n++ {
		nodes := make([]ListNode, n)
		for i := 0; i < n-1; i++ {
			nodes[i].Next = &nodes[i+1]
		}
		var head *ListNode
		if n > 0 {
			head = &nodes[0]
		}
		p := reverse(head)
		for i := n - 1; i >= 0; i-- {
			if p != &nodes[i] {
				t.Fatal("wrong identity/order", n, i)
			}
			p = p.Next
		}
		if p != nil {
			t.Fatal("cycle or extra node")
		}
	}
}
