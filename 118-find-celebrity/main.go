package main

func main() {}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		// space = O(1)
		candidate := 0
		// time = O(n), <= 3n
		for i := 1; i < n; i++ {
			if knows(candidate, i) {
				candidate = i
			}
		}
		for i := range n {
			if i != candidate {
				if knows(candidate, i) || !knows(i, candidate) {
					return -1
				}
			}
		}
		return candidate
	}
}
