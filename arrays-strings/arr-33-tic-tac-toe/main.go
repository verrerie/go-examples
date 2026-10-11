package main

import "slices"

func main() {}

type TicTacToe struct {
	cols  []int
	rows  []int
	diag1 int
	diag2 int
}

func Constructor(n int) TicTacToe {
	// space = O(n)
	cols := make([]int, n)
	rows := make([]int, n)
	return TicTacToe{
		cols: cols,
		rows: rows,
	}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	// time = O(1)
	n := len(t.rows)
	if player == 1 {
		t.rows[row]++
		t.cols[col]++
		if row == col {
			t.diag1++
		}
		if row+col == n-1 {
			t.diag2++
		}
	} else {
		t.rows[row]--
		t.cols[col]--
		if row == col {
			t.diag1--
		}
		if row+col == n-1 {
			t.diag2--
		}
	}
	reached := func(vars ...int) bool {
		return slices.Contains(vars, n) || slices.Contains(vars, -n)
	}
	if reached(t.rows[row], t.cols[col], t.diag1, t.diag2) {
		return player
	}
	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
