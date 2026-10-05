package main

import "fmt"

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}

func leastInterval(tasks []byte, n int) int {
	// space = O(26)
	counts := make(map[byte]int)
	highestFreq := 0
	// time = O(N)
	for _, task := range tasks {
		newFreq := counts[task] + 1
		counts[task] = newFreq
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
