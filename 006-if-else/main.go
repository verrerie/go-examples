package main

import "fmt"

func main() {
	n := 7

	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	n = 10
	d := 3

	if n%d == 0 {
		fmt.Println(n, "is divisible by", d)
	} else {
		fmt.Println(n, "is not divisible by", d)
	}

	if n%5 == 0 || d%5 == 0 {
		fmt.Println("eihter", n, "or", d, "is divisible by 5")
	}

	if n = 9; n < 0 {
		fmt.Println(n, "is negative")
	} else if n < 10 {
		fmt.Println(n, "has 1 digit")
	} else {
		fmt.Println(n, "has multiple digits")
	}
}
