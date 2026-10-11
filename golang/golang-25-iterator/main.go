package main

import (
	"fmt"
	"iter"
	"strings"
)

type element[T any] struct {
	val  T
	next *element[T]
}

type List[T any] struct {
	head, tail *element[T]
}

func (l *List[T]) Push(a T) {
	if l.tail == nil {
		l.head = &element[T]{
			a,
			nil,
		}
		l.tail = l.head
	} else {
		last := l.tail
		last.next = &element[T]{
			a,
			nil,
		}
		l.tail = last.next
	}
}

func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func fibo() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	l1 := List[int]{}
	l1.Push(10)
	l1.Push(20)
	l1.Push(30)
	fmt.Println(l1.head.val)
	fmt.Println(l1.tail.val)
	for a := range l1.All() {
		fmt.Println(a)
	}

	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Println(part)
	}

	for n := range fibo() {
		if n > 20 {
			return
		}
		fmt.Println(n)
	}
}
