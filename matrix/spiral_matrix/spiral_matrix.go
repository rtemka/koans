// #54
// https://leetcode.com/problems/spiral-matrix/description/
package leetcode

func spiralOrder(matrix [][]int) []int {
	var res []int
	// Initialize the matrix boundaries.
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	// Traverse the matrix in spiral order.
	for top <= bottom && left <= right {
		// Move from left to right along the top boundary.
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// Move from top to bottom along the right boundary.
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// Check that the bottom boundary hasn't passed the top boundary
		// before moving from right to left along the bottom boundary.
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		// Check that the left boundary hasn't passed the right boundary
		// before moving from bottom to top along the left boundary.
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
