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
	for x := 2; x <= n; x++ {
		s = rle(s)
	}
	return s
}

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
			fmt.Fprintf(&build, "%d%c", n, prev)
			n = 1
			prev = s[i]
		}
	}
	fmt.Fprintf(&build, "%d%s", n, string(prev))
	return build.String()
}
