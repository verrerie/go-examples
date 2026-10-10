package main

import "fmt"

func main() {
	fmt.Println(strStr("hello world!", "world"))
	fmt.Println(strStr("hello world!", "world!!"))
	fmt.Println(strStr("hello world!", "o w"))
}

// time = O((n−m+1) × m) = O(n x m)
// space = O(1)
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && haystack[i+j] == needle[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}
