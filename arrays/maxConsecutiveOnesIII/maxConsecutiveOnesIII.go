package leetcode

// Given a binary array nums and an integer k, return the maximum number of consecutive 1's
// in the array if you can flip at most k 0's.

// Example 1:

// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Example 2:
// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

// Constraints:

// 1 <= nums.length <= 105
// nums[i] is either 0 or 1.
// 0 <= k <= nums.length

func longestOnes(nums []int, k int) int {
	var max, ones, zeroes int
	var c = k

	for i, j := 0, 0; i < len(nums); i++ {

		if nums[i] == 0 && c == k {
			j = i
		}
		if zeroes == k {
			j = i - 1
			zeroes = 0
		}
		if nums[i] == 1 {
			ones++
			zeroes = 0
			if ones > max {
				max = ones
			}
		} else if c > 0 {
			ones++
			if ones > max {
				max = ones
			}
			c--
			zeroes++
		} else {
			c = k
			ones = 0
			i = j
			zeroes++
		}
		// log.Println("i: ", i, "j:", j, "ones:", ones, "max:", max, "z:", zeroes)
	}

	return max
}
