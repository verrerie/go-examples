package main

import "testing"

func TestMiddle(t *testing.T) {
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
		if middle(head) != want {
			t.Fatal(n)
		}
	}
}
