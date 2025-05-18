// #1385
// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/description/
package leetcode

import (
	"slices"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)
	res := 0
	for _, n := range arr1 {
		i, _ := slices.BinarySearch(arr2, n)
		if i < len(arr2) {
			if abs(n-arr2[i]) <= d {
				continue
			}
		}
		if i > 0 {
			if abs(n-arr2[i-1]) <= d {
				continue
			}
		}
		res++
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
