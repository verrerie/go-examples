package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle uses relative speed: within a cycle, fast gains one position per step.
// Time O(n); space O(1). Compare addresses, not values.
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func main() {
	h := &ListNode{Val: 1}
	h.Next = h
	fmt.Println(hasCycle(h))
}
