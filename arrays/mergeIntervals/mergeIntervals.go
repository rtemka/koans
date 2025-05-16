// #56
// https://leetcode.com/problems/merge-intervals/description/
package mergeintervals

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	lo := 0
	for hi := 1; hi < len(intervals); hi++ {
		if intervals[lo][1] >= intervals[hi][0] {
			intervals[lo][1] = max(intervals[lo][1], intervals[hi][1])
		} else {
			lo++
			intervals[lo] = intervals[hi]
		}
	}
	return intervals[:lo+1]
}
