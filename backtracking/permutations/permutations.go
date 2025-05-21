// #46
// https://leetcode.com/problems/permutations/description/
package leetcode

import "slices"

var ans [][]int

func permute(nums []int) [][]int {
	ans = nil
	memo := make(map[int]bool, len(nums))
	cur := make([]int, len(nums))
	backtrack(cur, nums, memo, 0)
	return ans
}

func backtrack(cur, nums []int, memo map[int]bool, idx int) {
	if len(nums) == idx {
		ans = append(ans, slices.Clone(cur))
		return
	}
	for i := range len(nums) {
		if memo[i] {
			continue
		}
		// Mark this number as used
		memo[i] = true
		// Place this number at current index
		cur[idx] = nums[i]
		// Recursively fill the next position
		backtrack(cur, nums, memo, idx+1)
		// // Backtrack: mark number as unused for next iteration
		memo[i] = false
	}
}
