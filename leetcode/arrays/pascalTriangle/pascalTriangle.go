package leetcode

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// Example 1:

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// Example 2:

// Input: numRows = 1
// Output: [[1]]

func generate(numRows int) [][]int {

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
