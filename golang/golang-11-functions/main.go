package main

import "fmt"

func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	plus := func(a, b int) int {
		return a + b
	}

	fmt.Println("a+b=", plus(1, 2))
	fmt.Println("a+b+c=", plusPlus(1, 2, 3))
}
