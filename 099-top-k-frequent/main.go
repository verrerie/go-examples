package main

import "fmt"

func main() {
	list := []int{1, 2, 1, 3, 1, 1, 2, 2}
	fmt.Println(topKFrequent(list, 2))
}

// time = O(n)
// space = O(n)
func topKFrequent(nums []int, k int) []int {
	// space: O(n)
	counts := map[int]int{}
	maxFreq := 0

	// O(n)
	for _, n := range nums {
		counts[n]++
		if counts[n] > maxFreq {
			maxFreq = counts[n]
		}
	}

	// O(c), c = n in worst case
	// space: O(n)
	buckets := make([][]int, maxFreq+1)
	for n, f := range counts {
		buckets[f] = append(buckets[f], n)
	}

	topK := []int{}
	// O(maxFreq + k), O(n) in worst case
	for k > len(topK) {
		topK = append(topK, buckets[maxFreq]...)
		maxFreq--
	}
	return topK
}
