package sorts

import "golang.org/x/exp/constraints"

func BubbleSort(arr []int) []int {
	// если мы проитерировали массив ни разу не меняя элементы,
	// то массив отсортирован
	var swaped bool

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[i] {
				arr[i], arr[j], swaped = arr[j], arr[i], true
			}

		}

		if !swaped {
			break
		}
	}

	return arr
}

func BubbleSortGeneric[T constraints.Ordered](arr []T) []T {
	// если мы проитерировали массив ни разу не меняя элементы,
	// то массив отсортирован
	var swaped bool

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[i] {
				arr[i], arr[j], swaped = arr[j], arr[i], true
			}

		}

		if !swaped {
			break
		}
	}

	return arr
}
