// #941
// https://leetcode.com/problems/valid-mountain-array/description/
package validmountainarray

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	peak := false
	for i, j, k := 0, 1, 2; k < len(arr); i, j, k = i+1, j+1, k+1 {
		switch {
		case (arr[i] < arr[j] && arr[j] < arr[k]), (arr[i] > arr[j] && arr[j] > arr[k]):
		case (arr[i] < arr[j] && arr[j] > arr[k]):
			peak = true
		default:
			return false
		}
	}
	return peak
}
