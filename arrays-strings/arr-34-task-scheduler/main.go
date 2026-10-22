// Question: Given a nonempty slice of tasks labeled A through Z and a nonnegative
// cooldown n, return the minimum number of time intervals needed to finish all
// tasks in any order. Each interval runs one task or is idle; equal task labels
// must have at least n intervening intervals.

package main

import "fmt"

func main() {
	fmt.Println(leastInterval([]byte("AAABBB"), 2))
	fmt.Println(leastInterval([]byte("ACABDB"), 1))
	fmt.Println(leastInterval([]byte("AAABBB"), 3))
}

func leastInterval(tasks []byte, n int) int {
	// space = O(26)
	counts := make(map[byte]int)
	highestFreq := 0
	// time = O(N)
	for _, task := range tasks {
		counts[task]++
		newFreq := counts[task]
		if newFreq > highestFreq {
			highestFreq = newFreq
		}
	}

	m := 0
	// time = O(1)
	for _, f := range counts {
		if f == highestFreq {
			m++
		}
	}

	return max(len(tasks), (highestFreq-1)*(n+1)+m)
}
