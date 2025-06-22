// #73
// https://leetcode.com/problems/set-matrix-zeroes/description
package leetcode

func setZeroes(matrix [][]int) {
	zrows := make(map[int]bool)
	zcols := make(map[int]bool)
	// Pass 1: Traverse through the matrix to identify the rows and
	// columns containing zeros and store their indexes in the appropriate hash sets.
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zrows[i] = true
				zcols[j] = true
			}
		}
	}
	// Pass 2: Set any cell in the matrix to zero if its row index is
	// in 'zrows' or its column index is in 'zcols'.
	for i := range matrix {
		for j := range matrix[i] {
			if zrows[i] || zcols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
