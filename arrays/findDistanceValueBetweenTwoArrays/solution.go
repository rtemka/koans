// #700
// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
package leetcode

import (
	"slices"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)
	res := 0
	for _, n := range arr1 {
		i, _ := slices.BinarySearch(arr2, n)
		i = min(len(arr2)-1, i)
		if abs(n-arr2[i]) > d {
			res++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
