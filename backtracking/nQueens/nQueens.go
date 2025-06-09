// #51
// https://leetcode.com/problems/n-queens/description/
package leetcode

import (
	"slices"
	"strings"
)

var (
	cols           []bool
	diag1          []bool
	diag2          []bool
	result         [][]string
	queenPositions []string
)

func solveNQueens(n int) [][]string {
	cols, diag1, diag2, result = make([]bool, n), make([]bool, n*2), make([]bool, n*2), nil
	board := make([]string, 0, n)
	queenPositions = make([]string, n)
	for i := range n {
		queenPositions[i] = strings.Repeat(".", i) + "Q" + strings.Repeat(".", n-i-1)
	}
	solve(0, n, board)
	return result
}

func solve(y, n int, board []string) {
	// If we succefully place all queens, save the result.
	if y == n {
		result = append(result, slices.Clone(board))
	}
	for x := 0; x < n; x++ {
		// If another queen "beats" this cell vertically(cols) or diagonally(diag 1 and 2), then skip.
		if cols[x] || diag1[x+y] || diag2[x-y+n-1] {
			continue
		}
		// Place the queen on board.
		cols[x], diag1[x+y], diag2[x-y+n-1] = true, true, true
		board = append(board, queenPositions[x])

		// Add another queen on horizontal line y+1.
		solve(y+1, n, board)

		// Remove queen from board.
		cols[x], diag1[x+y], diag2[x-y+n-1] = false, false, false
		board = board[:len(board)-1]
	}
}
