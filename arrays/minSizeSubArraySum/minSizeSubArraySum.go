package leetcode

import "math"

// Given an array of positive integers nums and a positive integer target, return the minimal length of a
// contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal
// to target. If there is no such subarray, return 0 instead.

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1
// Example 3:

// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

// Constraints:
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104

// Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

func minSubArrayLen(target int, nums []int) int {

	var sum, min int

	for i, j := 0, 0; min != 1; {

		if sum >= target {
			if j-i < min || min == 0 {
				min = j - i
			}
			sum -= nums[i]
			i++
		} else if j == len(nums) {
			break
		} else {
			sum += nums[j]
			j++
		}
		// log.Println("sum:", sum, "min:", min, "i:", i, "j:", j)
	}

	return min
}

func MinSubArrayLen2(target int, nums []int) int {

	var sum, lo, hi = 0, 0, 0
	var min = math.MaxInt

	for hi < len(nums) {

		sum += nums[hi]

		for target <= sum {
			if hi-lo+1 < min {
				min = hi - lo + 1
			}
			sum -= nums[lo]
			lo++
		}
		hi++
	}

	if min == math.MaxInt {
		return 0
	}
	return min
}
