package leetcode

// Целочисленный массив называется арифметическим, если он состоит как минимум из трех элементов
// и если разница между любыми двумя последовательными элементами одинакова.

// Например, [1,3,5,7,9], [7,7,7,7] и [3,-1,-5,-9] являются арифметическими последовательностями.
// Учитывая целочисленный массив nums, вернуть количество арифметических подмассивов nums.

// Подмассив — это непрерывная подпоследовательность массива.

// Пример 1:

// Ввод: nums = [1,2,3,4]
// Выход: 3
// Объяснение: у нас есть 3 арифметических среза в nums: [1, 2, 3], [2, 3, 4] и собственно [1,2,3,4].

// Пример 2:

// Ввод: nums = [1]
// Выход: 0

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	return countValidSubArrays(nums, 3, isArith)
}

func isArith(arr []int, n int) bool {
	if len(arr) < n || n < 3 {
		return false
	}
	// берём первый, второй и последний элементы
	f, s, l := arr[0], arr[1], arr[len(arr)-1]
	var diff = diff(f, s)
	if f > l {
		// предполагаемое значене последнего элемента
		pl := f - ((len(arr) - 1) * diff)
		// проверяем соблюдался ли шаг между элементами
		if pl != l {
			return false
		}
	} else {
		// предполагаемое значение первого элемента
		pf := l - ((len(arr) - 1) * diff)
		// проверяем соблюдался ли шаг между элементами
		if pf != f {
			return false
		}
	}

	return true
}

func countValidSubArrays(arr []int, n int, f func([]int, int) bool) int {
	var c int
	for i := n; i <= len(arr); i++ {
		for j := 0; j+i <= len(arr); j++ {
			if f(arr[j:i+j], i) {
				c++
			}
		}
	}
	return c
}

func diff(f, s int) int {
	if (f < 0 && s > 0) || (s < 0 && f > 0) {
		return abs(f) + abs(s) // -a1, a2 ... n || a1, -a2 ... n
	} else {
		as, af := abs(s), abs(f)
		if as > af {
			return as - af // -a1, -a2, a3 ... n || a1, a2, -a3 ... n
		} else {
			return af - as // a1, a2 ... n || -a1, -a2 ... n
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
