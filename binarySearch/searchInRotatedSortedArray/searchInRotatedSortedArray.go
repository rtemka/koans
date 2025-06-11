// #33
// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
package leetcode

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		// One half will be sorted, so determine that sorted half.
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
			// Otherwise the right half is sorted.
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
