package main

import "fmt"

func main() {
	fmt.Println("slice:", []int{1, 2, 3})

	var a []int
	fmt.Println("a:", a)
	fmt.Println("len:", len(a))
}
