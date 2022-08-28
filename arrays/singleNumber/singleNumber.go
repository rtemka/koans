package leetcode

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Example 1:

// Input: nums = [2,2,1]
// Output: 1
// Example 2:

// Input: nums = [4,1,2,1,2]
// Output: 4
// Example 3:

// Input: nums = [1]
// Output: 1

func singleNumber(nums []int) int {
	n := 0

	for _, v := range nums {
		n = n ^ v
	}

	return n
}

func singleNumberNaive(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
