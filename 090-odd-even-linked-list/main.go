package main

import (
	"fmt"
	"slices"
)

func main() {
	l := makeList(1, 2, 3, 4, 5, 6)
	show(oddEvenList(l))
	l = makeList(1, 2, 3, 4, 5, 6, 7)
	show(oddEvenList(l))
}

func makeList(nums ...int) *ListNode {
	var head *ListNode
	for _, num := range slices.Backward(nums) {
		temp := head
		head = &ListNode{Val: num}
		head.Next = temp
	}
	return head
}

func show(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		fmt.Print("->")
		l = l.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// space = O(1)
	oddHead := head
	evenHead := head.Next
	odd := oddHead
	even := evenHead
	// O(n)
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = evenHead
	return oddHead
}
