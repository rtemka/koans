package search

func BinaryIndex(arr []int, x int) int {

	for start, end := 0, len(arr)-1; end > start; {

		mid := (end + start) / 2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] < x {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1
}

func BinaryHas(arr []int, x int) bool {
	l := arr[:]

	for mid := len(l) / 2; len(l) > 0; mid = len(l) / 2 {

		if l[mid] == x {
			return true
		}

		if l[mid] > x {
			l = l[:mid]
		} else {
			l = l[mid+1:]
		}
	}
	return false
}
