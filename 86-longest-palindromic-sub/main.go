package main

import "fmt"

func main() {
	fmt.Println(expandAndFind("a", 0, 0))
	fmt.Println(expandAndFind("aa", 0, 1))
	fmt.Println(expandAndFind("", 0, 0))
	fmt.Println(expandAndFind("aba", 0, 0))
	fmt.Println(expandAndFind("aba", 1, 1))

	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	maxLength := 0
	result := ""
	for i := range len(s) {
		length, palindrom := expandAndFind(s, i, i)
		if length > maxLength {
			maxLength = length
			result = palindrom
		}
		length, palindrom = expandAndFind(s, i, i+1)
		if length > maxLength {
			maxLength = length
			result = palindrom
		}

	}
	return result
}

func expandAndFind(s string, l, r int) (int, string) {
	length := 0
	result := ""
	for l >= 0 && r < len(s) && s[l] == s[r] {
		//	when valid we expand
		length = r - l + 1
		result = s[l : r+1]
		l--
		r++
	}
	return length, result
}
