// #59
// https://leetcode.com/problems/spiral-matrix-ii/description/
package leetcode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, n)
	}
	x := 1
	i, n := 0, n-1
	for i < n {
		h, v := i, i
		for h < n {
			matrix[v][h] = x
			x++
			h++
		}
		for v < n {
			matrix[v][h] = x
			x++
			v++
		}
		for h > i {
			matrix[v][h] = x
			x++
			h--
		}
		for v > i {
			matrix[v][h] = x
			x++
			v--
		}
		i++
		n--
	}
	if matrix[i][i] == 0 {
		matrix[i][i] = x
	}
	// fmt.Println(matrix)
	return matrix
}

