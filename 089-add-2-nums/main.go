package main

import (
	"fmt"
	"slices"
)

func main() {
	// 1 -> 4 -> 5
	// 9 -> 5 -> 4
	// = 0 -> 0 -> 0 -> 1

	// nil
	// nil
	// = 0

	// 1
	// 9
	// = 0 -> 1
	l1 := makeList(1, 2, 3, 4)
	show(l1)
	l2 := makeList(3, 5, 6, 0)
	show(l2)
	ll := addTwoNumbers(l1, l2)
	show(ll)

	l1 = makeList(1, 4, 5)
	l2 = makeList(9, 5, 4)
	ll = addTwoNumbers(l1, l2)
	show(ll)

	show(addTwoNumbers(nil, nil))
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

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	n1 := 0
	n2 := 0
	carry := 0
	//	O(max(L1,L2), extra space = O(1), exlucding the result
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		sum := n1 + n2 + carry
		tail.Val = sum % 10
		carry = sum / 10
		if l1 != nil || l2 != nil || carry > 0 {
			tail.Next = &ListNode{}
			tail = tail.Next
		} else {
			tail.Next = nil // last one
		}

	}
	return head
}
