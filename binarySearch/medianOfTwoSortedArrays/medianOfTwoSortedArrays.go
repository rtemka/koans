// #4
// https://leetcode.com/problems/jump-game/description/
// Editorial: https://leetcode.com/problems/median-of-two-sorted-arrays/editorial/#approach-3-a-better-binary-search
package leetcode

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	var L1, L2, R1, R2 int
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	for {
		L1Index := (left + right) / 2
		L2Index := (m+n+1)/2 - L1Index

		if L1Index == 0 {
			L1 = math.MinInt64
		} else {
			L1 = nums1[L1Index-1]
		}
		if L1Index == m {
			R1 = math.MaxInt64
		} else {
			R1 = nums1[L1Index]
		}
		if L2Index == 0 {
			L2 = math.MinInt64
		} else {
			L2 = nums2[L2Index-1]
		}
		if L2Index == n {
			R2 = math.MaxInt64
		} else {
			R2 = nums2[L2Index]
		}
		if L1 > R2 {
			right = L1Index - 1
		} else if L2 > R1 {
			left = L1Index + 1
		} else {
			if (m+n)%2 == 0 {
				maxLeft := max(L1, L2)
				minRight := min(R1, R2)
				return float64(maxLeft+minRight) / 2.0
			} else {
				return float64(max(L1, L2))
			}
		}
	}
}
