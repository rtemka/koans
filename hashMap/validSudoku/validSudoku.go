// #36
// https://leetcode.com/problems/valid-sudoku/description/
package leetcode

func isValidSudoku(board [][]byte) bool {
	// Create array of 9 rows to keep track previously seen numbers in each row.
	rows := [9][9]byte{}
	// Create array of 9 columns to keep track previously seen numbers in each column.
	cols := [9][9]byte{}
	// Create table 3x3 to keep track previously seen numbers in each of 9 subgrids in sudoku board.
	subgrid := [3][3][9]byte{}

	for r := range 9 {
		for c := range 9 {
			n := board[r][c]
			// Skip the blank cells.
			if n == '.' {
				continue
			}
			// Calculate the index in the 1..9 range.
			i := n - '1'
			// Check if 'num' has been seen in the current row, column, or subgrid.
			if n == rows[r][i] ||
				n == cols[c][i] ||
				n == subgrid[r/3][c/3][i] {
				return false
			}
			// If check is passed, save this number in row, column and subgrid arrays.
			rows[r][i] = n
			cols[c][i] = n
			subgrid[r/3][c/3][i] = n
		}
	}
	return true
}
