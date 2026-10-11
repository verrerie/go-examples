package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"abc", "edf", "aa"}

	ints := []int{4, 8, 1}

	slices.Sort(strs)
	slices.Sort(ints)

	fmt.Println(strs)
	fmt.Println(ints)
	fmt.Println(slices.IsSorted(strs))
	fmt.Println(slices.IsSorted(ints))
}
