package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	collection := map[[26]int][]string{}
	// O(N)
	for _, word := range strs {
		counts := [26]int{}
		// O(L)
		for _, ch := range word {
			counts[ch-'a']++
		}
		collection[counts] = append(collection[counts], word)
	}
	for _, list := range collection {
		result = append(result, list)
	}
	return result
}
