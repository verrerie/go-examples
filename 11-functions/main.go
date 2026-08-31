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

	concat := func(ss ...string) string {
		result := ""
		for _, s := range ss {
			result = result + s
		}
		return result
	}
	fmt.Println("concated:", concat("1", "22", "333", "4444", "end"))

	initSeq := func(init int) func() int {
		i := init
		return func() int {
			i = i + 1
			return i
		}
	}

	inc := initSeq(5)
	fmt.Println("value=", inc())
	fmt.Println("value=", inc())
	fmt.Println("value=", inc())

	next := initSeq(0)
	fmt.Println("newValue=", next())

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
	fmt.Println("c(6,3)=", c(3, 2))
}
