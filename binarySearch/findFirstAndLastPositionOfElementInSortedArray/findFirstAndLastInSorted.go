// #34
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
package leetcode

import "slices"

func searchRange(nums []int, target int) []int {
	// Implementation of BinarySearch in Go returns the
	// leftmost position of the target when target have duplicates.
	i, ok := slices.BinarySearch(nums, target)
	if ok {
		// If we found target then search for the leftmost target+1.
		j, _ := slices.BinarySearch(nums, target+1)
		// If the target+1 is found: j-1 gives us the rightmost target index.
		// If the target+1 is not found: j == len(nums); j-1 gives us the rightmost target index.
		// So j-1 gives us the rightmost target index anyway.
		return []int{i, j - 1}
	} else {
		return []int{-1, -1}
	}
}
