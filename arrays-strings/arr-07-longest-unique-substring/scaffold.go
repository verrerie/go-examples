package main

// Original sliding-window teaching scaffold. lengthOfLongestSubstring is the implementation.
const WINDOW_CONDITION_BROKEN = false

func fn3(arr []int) int {
	var left, ans, curr int

	for right := range len(arr) {
		// do logic to add arr[right] to curr
		curr = curr + arr[right]
		for WINDOW_CONDITION_BROKEN {
			// remove arr[left] from curr
			curr = curr - arr[left]
			left++
		}

		// update ans
		ans = ans + curr
	}

	return ans
}
