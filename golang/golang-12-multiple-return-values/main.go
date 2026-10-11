package main

import "fmt"

func main() {
	div := func(a, b int) (bool, int) {
		if b == 0 {
			return false, 0
		}
		return true, a / b
	}

	success, v := div(20, 5)
	fmt.Println("a / b =", success, ",", v)
	success, v = div(20, 0)
	fmt.Println("a / 0 =", success, ",", v)
}
