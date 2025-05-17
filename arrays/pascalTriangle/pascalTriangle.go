// #118
// https://leetcode.com/problems/pascals-triangle/description/
package leetcode

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	res = append(res, []int{1, 1})

	for i := 1; i < numRows-1; i++ {
		row := make([]int, 0, len(res[i])+1)
		// Add one to start.
		row = append(row, 1)
		for k := 1; k < len(res[i])-1; k++ {
			row = append(row, res[i][k]+res[i][k+1])
		}
		// Add one to the end
		row = append(row, 1)
		res = append(res, row)
	}

	return res
}

func GenerateOld(numRows int) [][]int {

	if numRows == 0 {
		return [][]int{}
	}

	t := make([][]int, 0, numRows)
	t = append(t, []int{1}) // добавляем корень

	for i, p := 1, 0; i < numRows; i, p = i+1, p+1 {

		var r = make([]int, i+1) // текущий ряд
		f, s := 0, 1             // указатели на первый и второй элементы предыдущего ряда

		for j := range r {
			// если второй указатель не ушел за границы пред. ряда
			// и если мы не на первом или последнем элементе текущего ряда
			if s < i+1 && j != 0 && j != len(r)-1 {
				r[j] = t[p][f] + t[p][s] // добавляем сумму первого и второго элемента
				f, s = f+1, s+1          // и двигаем указатели
			} else {
				r[j] = t[p][f] // добавляем указатель на первый/последний элемент
				// и не двигаем указатели
			}
		}

		t = append(t, r) // добавляем ряд в массив
	}

	return t
}
