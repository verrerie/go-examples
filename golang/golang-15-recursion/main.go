package main

import "fmt"

func main() {
	var c func(n, k int) int
	c = func(n, k int) int {
		if k == 1 {
			return n
		}
		if k > n {
			return -1
		}
		return c(n-1, k-1) * n / k
	}
	fmt.Println("c(3,2)=", c(3, 2))
	fmt.Println("c(6,2)=", c(6, 2))
	fmt.Println("c(6,3)=", c(6, 3))
}
