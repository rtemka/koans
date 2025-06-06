// #57
// https://leetcode.com/problems/insert-interval/description/
package leetcode

import (
	"slices"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	cmpfunc := func(a []int, b int) int {
		if b > a[1] {
			return -1
		} else if b < a[0] {
			return +1
		} else {
			return 0
		}
	}

	lo, ok := slices.BinarySearchFunc(intervals, newInterval[0], cmpfunc)
	if ok {
		newInterval[0] = min(intervals[lo][0], newInterval[0])
	}
	hi, ok := slices.BinarySearchFunc(intervals, newInterval[1], cmpfunc)
	if ok {
		newInterval[1] = max(intervals[hi][1], newInterval[1])
		hi++
	}

	res := append([][]int{}, intervals[:lo]...)
	res = append(res, newInterval)
	res = append(res, intervals[hi:]...)

	return res
}
