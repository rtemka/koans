package leetcode

// Given an array nums with n objects colored red, white, or blue,
// sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

// You must solve this problem without using the library's sort function.

// Example 1:
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]

// Example 2:
// Input: nums = [2,0,1]
// Output: [0,1,2]

// Constraints:
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.

// Follow up: Could you come up with a one-pass algorithm using only constant extra space?

// [0,0,1,1,2,2]
//          ^
//            ^

// решение методом двух указателей
func sortColors(nums []int) {
	var min int
	for i, j := 0, 1; min < 2 && j < len(nums); j++ {
		if nums[j] < nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i] == min {
			i++
		}
		if j == len(nums)-1 {
			min++
			j = i
		}
	}
}
