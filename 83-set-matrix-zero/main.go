package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

// space: O(1)
func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m <= 0 {
		return
	}
	n := len(matrix[0])
	firstRowZeros := false
	firstColZeros := false

	for i := range m {
		if matrix[i][0] == 0 {
			firstColZeros = true
			break
		}
	}

	for j := range n {
		if matrix[0][j] == 0 {
			firstRowZeros = true
			break
		}
	}

	// O(M * N)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// O(M * N)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZeros {
		for j := range n {
			matrix[0][j] = 0
		}
	}

	if firstColZeros {
		for i := range m {
			matrix[i][0] = 0
		}
	}
}
