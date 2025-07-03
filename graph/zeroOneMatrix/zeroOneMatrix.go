// #541
// https://leetcode.com/problems/01-matrix/description/
package leetcode

import "container/list"

// Define direction vectors for up, down, left and right.
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Position struct {
	row int
	col int
}

func updateMatrix(mat [][]int) [][]int {
	// Queue for the BFS.
	queue := list.New()
	// Memo for the visited cells.
	visited := make([][]int, len(mat))
	for i := range visited {
		visited[i] = make([]int, len(mat[0]))
	}
	// Gather all '0' cells from the matrix.
	for row := range mat {
		for col := range mat[row] {
			if mat[row][col] == 0 {
				queue.PushBack(Position{row: row, col: col})
				visited[row][col] = 1
			}
		}
	}

	// Perfom multi-source BFS from '0' cells continuously updating
	// adjacent '1' cells with the shortest distance to '0'.
	distance := 0
	for queue.Len() > 0 {
		// Add 1 to the distance with each level of the matrix that's explored.
		distance++
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(Position)
			// Update any neighboring cells and add them to the queue to be processed in
			// the next level.
			for _, dir := range dirs {
				next_r, next_c := node.row+dir[0], node.col+dir[1]
				if isWithinBounds(next_r, next_c, mat) && visited[next_r][next_c] == 0 && mat[next_r][next_c] == 1 {
					mat[next_r][next_c] = distance
					visited[next_r][next_c] = 1
					queue.PushBack(Position{row: next_r, col: next_c})
				}
			}
		}
	}

	return mat
}

func isWithinBounds(row, col int, matrix [][]int) bool {
	return 0 <= row && row < len(matrix) && 0 <= col && col < len(matrix[0])
}
