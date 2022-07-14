package sorts

import "golang.org/x/exp/constraints"

// Найдите позицию min наименьшего элемента в диапазоне [i...n−1],
// Поменяйте элемент на i-й элемент,
// Увеличьте нижнюю границу i на 1 и повторяйте шаг 1, пока i = N-2.
func SelectionSort(arr []int) []int {

	for i, min := 0, 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if arr[min] != arr[i] {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}

// Найдите позицию min наименьшего элемента в диапазоне [i...n−1],
// Поменяйте элемент на i-й элемент,
// Увеличьте нижнюю границу i на 1 и повторяйте шаг 1, пока i = N-2.
func SelectionSortGeneric[T constraints.Ordered](arr []T) []T {

	for i, min := 0, 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if arr[min] != arr[i] {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	return arr
}
