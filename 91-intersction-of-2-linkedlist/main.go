package main

import (
	"fmt"
	"slices"
)

func main() {
	l1 := makeList(1, 2, 3, 4)
	fmt.Println(length(l1))
	l2 := makeList(5, 6, 7, 8)
	fmt.Println(length(l2))
	fmt.Println(getIntersectionNode(l1, l2))
	fmt.Println(getIntersectionNode(l1, nil))
	fmt.Println(getIntersectionNode(nil, nil))
	tail1 := l1
	for tail1.Next != nil {
		tail1 = tail1.Next
	}
	tail1.Next = l2
	fmt.Println(*getIntersectionNode(l1, l2))
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

func length(l *ListNode) int {
	result := 0
	for l != nil {
		result++
		l = l.Next
	}
	return result
}

func forward(l *ListNode, n int) *ListNode {
	newHead := l
	for n > 0 && newHead != nil {
		n--
		newHead = newHead.Next
	}
	return newHead
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// O(m)
	lenA := length(headA)
	// O(n)
	lenB := length(headB)
	diff := lenA - lenB
	if diff > 0 {
		headA = forward(headA, diff)
	} else if diff < 0 {
		headB = forward(headB, -diff)
	}
	// scan the common node
	// O(min(m,n))
	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	// space = O(1)
	return headA
}
