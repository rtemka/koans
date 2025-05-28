package sorts2025

func QuickSort(s []int) []int {
	return quicksort(s, 0, len(s)-1)
}

func quicksort(s []int, lo, hi int) []int {
	for lo < hi {
		var p int
		s, p = partition(s, lo, hi)
		s = quicksort(s, lo, p-1)
		s = quicksort(s, p+1, hi)
	}
	return s
}

func partition(s []int, lo, hi int) ([]int, int) {
	pivot := s[hi]
	for j := lo; j < hi; j++ {
		if s[j] < pivot {
			s[lo], s[j] = s[j], s[lo]
			lo++
		}
	}
	s[lo], s[hi] = s[hi], s[lo]
	return s, lo
}
