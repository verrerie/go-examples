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
	// we need to expand from every center from s
	// time <= O(n * (2n-1)) = O(n*n)
	// space <= O(1)
	for i := range len(s) {
		// time <= O(n)
		length, palindrome := expandAndFind(s, i, i)
		if length > maxLength {
			maxLength = length
			result = palindrome
		}
		// time <= O(n-1)
		length, palindrome = expandAndFind(s, i, i+1)
		if length > maxLength {
			maxLength = length
			result = palindrome
		}

	}
	return result
}

// time = O(min(l+1, n-r)), we expand outward from the center until a mismatch or boundary
// worst case ~n/2 steps (center of "aaa...a") <= O(n)
func expandAndFind(s string, l, r int) (int, string) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		//	when valid we expand
		l--
		r++
	}
	// the loop stopped one step past the palindrome, which is s[l+1:r]
	length := r - l - 1
	if length <= 0 {
		return 0, ""
	}
	// time = O(1), extra space = O(1) as the substring shares the same memory of parent s
	return length, s[l+1 : r]
}
