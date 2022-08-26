package sorts

import "sort"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElementNaive(nums []int) int {

	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	max, maj := 0, 0
	for k, v := range m {
		if v > max {
			max, maj = v, k
		}
	}

	return maj
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
