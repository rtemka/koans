package leetcode

// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:

// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:

// Input: nums = [1,3,5,6], target = 7
// Output: 4

func searchInsert(nums []int, target int) int {

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

func SearchInsert(nums []int, target int) int {

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2

		diff := target - nums[mid]
		if diff < 0 {
			hi = mid - 1
		} else if diff > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}

	if target > nums[lo] {
		return lo + 1
	} else {
		return lo
	}
}
