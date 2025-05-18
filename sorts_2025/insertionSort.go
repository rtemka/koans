package sorts2025

func InsertionSort(arr []int) []int {
	for i := range len(arr) {
		for j := i + 1; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}
