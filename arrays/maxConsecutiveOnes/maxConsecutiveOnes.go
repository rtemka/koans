package leetcode

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Example 1:

// Input: nums = [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
// Example 2:

// Input: nums = [1,0,1,1,0,1]
// Output: 2

// Constraints:

// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.

// метод скользящего окна с двумя указателями
func findMaxConsecutiveOnes(nums []int) int {
	var max, i, j int

	for ; j < len(nums); j++ {

		if nums[j] == 0 {
			c := j - i
			if c > max {
				max = c
			}
			i = j + 1
		}

	}

	if j-i > max {
		return j - i
	} else {
		return max
	}

}

// метод простого подсчета
func findMaxConsecutiveOnesByCount(nums []int) int {
	var max, count int

	for i := range nums {
		if nums[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}

	return max
}
