package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(3, 2))
	fmt.Println(getSum(-1, 2))
	fmt.Println(getSum(-1, -2))
}

func getSum(a int, b int) int {
	for b != 0 {
		sumWithoutCarry := a ^ b
		carry := (a & b) << 1
		a, b = sumWithoutCarry, carry
	}
	return a
}
