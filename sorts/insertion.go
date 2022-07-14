package sorts

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](arr []T) []T {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
		arr[j+1] = key
	}

	return arr
}
