// #200
// https://leetcode.com/problems/number-of-islands/description
package leetcode

import "container/list"

// Define direction vectors for up, down, left and right.
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	m, n := len(grid), len(grid[0])
	for row := range m {
		for col := range n {
			// If a land cell is found, perform DFS to explore the full
			// island, and include this island in our count.
			if grid[row][col] == '1' {
				DFS(row, col, grid)
				count++
			}
		}
	}
	return count
}

func DFS(row, col int, grid [][]byte) {
	// Mark the current cell as visited.
	grid[row][col] = '#'
	// Recursively call DFS on each neighboring land cell to continue
	// exploring this island.
	for _, d := range dirs {
		nextRow, nextCol := row+d[0], col+d[1]
		if isWithinBounds(nextRow, nextCol, grid) && grid[nextRow][nextCol] == '1' {
			DFS(nextRow, nextCol, grid)
		}
	}
}

// Chech that provided row and column is within grid.
func isWithinBounds(row, col int, grid [][]byte) bool {
	return 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0])
}

type Position struct {
	row int
	col int
}

func BFS(row, col int, grid [][]byte) {
	// Make a queue.
	queue := list.New()
	// Mark the current cell as visited.
	grid[row][col] = '#'

	queue.PushBack(Position{row: row, col: col})
	for queue.Len() > 0 {
		// Pop node from the front of the queue.
		node := queue.Remove(queue.Front()).(Position)
		for _, d := range dirs {
			nextRow, nextCol := node.row+d[0], node.col+d[1]
			// Add each neighboring land cell to continue
			// exploring this island.
			if isWithinBounds(nextRow, nextCol, grid) && grid[nextRow][nextCol] == '1' {
				queue.PushBack(Position{row: nextRow, col: nextCol})
				grid[nextRow][nextCol] = '#'
			}
		}
	}
}
