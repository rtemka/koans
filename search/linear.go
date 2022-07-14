package search

func LinearSearchIndex(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}

	return -1
}
