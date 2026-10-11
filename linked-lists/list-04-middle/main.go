package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// middle returns the second middle on an even-length acyclic list.
// Invariant: after t steps, slow moved t links and fast moved 2t links.
// Time O(n); space O(1).
func middle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(middle(h).Val)
}
