// #74
// https://leetcode.com/problems/search-a-2d-matrix/description/
package leetcode

// import "slices"

func searchMatrix(matrix [][]int, target int) bool {
	lo, hi := 0, len(matrix)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// _, ok := slices.BinarySearch(matrix[mid], target)
		if _, ok := binarySearch(matrix[mid], target); ok {
			return true
		}
		if target > matrix[mid][len(matrix[mid])-1] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func binarySearch(s []int, target int) (idx int, ok bool) {
	n := len(s)
	lo, hi := 0, n
	for lo < hi {
		mid := (lo + hi) / 2
		if s[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo, lo < n && s[lo] == target
}
