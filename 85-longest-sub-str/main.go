package main

import "fmt"

func main() {
	input := "abdebecf"
	fmt.Println("input:", input)
	fmt.Println("result:", lengthOfLongestSubstring(input))
	fmt.Println(lengthOfLongestSubstring("abcdfe"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("a"))
}

func lengthOfLongestSubstring(s string) int {
	counts := make(map[byte]int)
	result := 0 // empty string return 0
	i := 0
	// we make sure every step / iteration of j, the window between i and j is valid
	// time = O(n)
	for j := 0; j < len(s); j++ {
		counts[s[j]] = counts[s[j]] + 1
		// remove duplicates: we have only one char that has only one duplicate
		// so we keep moving forward until the duplicate is removed
		for counts[s[j]] > 1 {
			counts[s[i]] = counts[s[i]] - 1
			i++ // advance i
		}
		// now the window is valid (without duplicated between i and j)
		result = max(result, j-i+1)
	}

	return result
}
