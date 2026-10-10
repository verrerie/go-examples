package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// time = O(K x L)
// space = O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	j := 0

	for j < len(strs[0]) {
		for _, str := range strs {
			if j == len(str) || strs[0][j] != str[j] {
				return strs[0][:j]
			}
		}
		j++
	}
	return strs[0][:j]
}
