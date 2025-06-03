// #35
// https://leetcode.com/problems/search-insert-position/description/
package leetcode

func searchInsert(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func searchInsertOld(nums []int, target int) int {

	start, end, mid := 0, len(nums), 0

	for end > start {
		mid = (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid
		}

	}

	if target > nums[mid] {
		return mid + 1
	} else {
		return mid
	}
}
