package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rle("111221"))
	fmt.Println(rle("1111"))
	fmt.Println(rle("1"))
	fmt.Println(rle(""))
	fmt.Println(rle("123454"))

	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	s := "1"
	// time = O(L_1 + ... + L_n), L_k = len of k-th term ~ λ^k, λ ≈ 1.3036 (Conway's constant)
	// geometric series, so O(λ^n)
	// space = O(λ^n), the final string
	for x := 2; x <= n; x++ {
		s = rle(s)
	}
	return s
}

// time = O(L), space = O(L), L = len(s)
// counts are written as a single digit, which holds for count-and-say terms (runs <= 3) but not for arbitrary input
func rle(s string) string {
	if len(s) == 0 {
		return s
	}
	build := strings.Builder{}
	prev := s[0]
	n := 0
	for i := 0; i < len(s); i++ {
		if prev == s[i] {
			n++
		} else {
			// the ascii number of 1 single digit
			build.WriteByte('0' + byte(n))
			build.WriteByte(prev)
			n = 1
			prev = s[i]
		}
	}
	build.WriteByte('0' + byte(n))
	build.WriteByte(prev)
	return build.String()
}
