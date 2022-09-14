package leetcode

import (
	"log"
	"math"
)

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// You must do it in place.

// Example 1:
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]

// Example 2:
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Constraints:
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2^31 <= matrix[i][j] <= 2^31 - 1

// Follow up:
// A straightforward solution using O(mn) space is probably a bad idea.
// A simple improvement uses O(m + n) space, but still not the best solution.
// Could you devise a constant space solution?

// [t,t,f]
// [1,1,1][f]
// [0,0,1][t]
// [1,1,1][f]

// [t,t,f]
// [0,0,1][f]
// [0,0,0][t]
// [0,0,1][f]

// с привлечение доп. памяти M+N
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var rows = make([]bool, len(matrix))
	var cols = make([]bool, len(matrix[0]))

	for i := 0; i < len(cols); i++ {
		for j := 0; j < len(rows); j++ {
			if matrix[j][i] == 0 {
				cols[i] = true
				rows[j] = true
			}
		}
	}
	for i := range cols {
		if cols[i] {
			for j := 0; j < len(rows); j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := range rows {
		if rows[i] {
			for j := 0; j < len(cols); j++ {
				matrix[i][j] = 0
			}
		}
	}
}

// без привлечения доп. памяти M+N
func setZeroesConstSapce(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] == 0 {
				if matrix[j][i] == math.MinInt {
					matrix[j][i] = 0
				}
				matrix[0][i] = math.MaxInt
				matrix[j][0] = math.MinInt
			}
		}
	}
	log.Println(matrix)
}
