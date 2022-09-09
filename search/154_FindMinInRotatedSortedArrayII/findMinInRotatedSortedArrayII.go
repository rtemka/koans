package leetcode

import (
	"math"
)

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// For example, the array nums = [0,1,4,4,5,6,7] might become:

// [4,5,6,7,0,1,4] if it was rotated 4 times.
// [0,1,4,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

// Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.

// You must decrease the overall operation steps as much as possible.

// Example 1:
// Input: nums = [1,3,5]
// Output: 1

// Example 2:
// Input: nums = [2,2,2,0,1]
// Output: 0

// Example 3:
// Input: nums = [3,3,3,3,1,3]
// Output: 1

// Constraints:
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// nums is sorted and rotated between 1 and n times.

func findMin(nums []int) int {
	var min = math.MaxInt
	var lo = 0
	var hi = len(nums) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] < min {
			min = nums[mid]
		}
		if lo == hi {
			return min
		}
		if nums[mid] == nums[lo] && nums[mid] == nums[hi] {
			ml := findMin(nums[lo:mid])
			mh := findMin(nums[mid:hi])
			if ml < min {
				min = ml
			}
			if mh < min {
				min = mh
			}
			return min
		}
		if nums[mid] <= nums[hi] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return min
}
