package main

import "fmt"

func main() {
	grid := [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(grid)) // Expected: 3
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	h, w := len(grid), len(grid[0])
	// space = O(q), q = (w * h) in worst case
	queue := make([][2]int, 0)
	count := 0
	// time = O(w * h)
	for i := range h {
		for j := range w {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = 'x'
				queue = append(queue, [2]int{i, j})

				for len(queue) > 0 {
					next := queue[0]
					queue = queue[1:]
					r, c := next[0]+1, next[1]
					if r < h && grid[r][c] == '1' {
						grid[r][c] = 'x'
						queue = append(queue, [2]int{r, c})
					}
					r, c = next[0]-1, next[1]
					if r >= 0 && grid[r][c] == '1' {
						grid[r][c] = 'x'
						queue = append(queue, [2]int{r, c})
					}
					r, c = next[0], next[1]+1
					if c < w && grid[r][c] == '1' {
						grid[r][c] = 'x'
						queue = append(queue, [2]int{r, c})
					}
					r, c = next[0], next[1]-1
					if c >= 0 && grid[r][c] == '1' {
						grid[r][c] = 'x'
						queue = append(queue, [2]int{r, c})
					}
				}
			}
		}
	}
	return count
}
