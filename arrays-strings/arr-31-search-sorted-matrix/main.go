package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(searchMatrix(matrix, 25))
	fmt.Println(searchMatrix(matrix, 30))

	matrix = [][]int{
		{1, 1},
	}
	fmt.Println(searchMatrix(matrix, 0))
}

// time = O(h + w), space = O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	h := len(matrix)
	w := len(matrix[0])

	// start from top right
	c, r := w-1, 0
	// time = O(h + w)
	for c >= 0 && r < h {
		switch {
		case matrix[r][c] > target:
			c--
		case matrix[r][c] < target:
			r++
		default:
			// found
			return true
		}
	}
	return false
}
