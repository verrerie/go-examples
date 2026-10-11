package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverse changes an acyclic list in place.
// Invariant: prev is the reversed prefix; curr starts the untouched suffix.
// Time O(n); space O(1).
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // Save the suffix before replacing the link.
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
func main() {
	for p := reverse(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}); p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
