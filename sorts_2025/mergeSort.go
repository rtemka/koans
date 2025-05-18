package sorts2025

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	for ; l < len(left); l++ {
		result = append(result, left[l])
	}
	for ; r < len(right); r++ {
		result = append(result, right[r])
	}
	return result
}
