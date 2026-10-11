package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "orange", "kiwi"}

	lenComp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenComp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	pers := []Person{
		{"bob", 41},
		{"alice", 33},
		{"thomas", 34},
	}

	slices.SortFunc(pers, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(pers)
}
