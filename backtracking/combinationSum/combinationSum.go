// #39
// https://leetcode.com/problems/combination-sum/description/
package leetcode

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var backtrack func(comb []int, start, target int)

	backtrack = func(comb []int, start, target int) {
		// If the target is equal to 0, we found a combination that sums to initial target.
		if target == 0 {
			res = append(res, slices.Clone(comb))
			return
		}
		// If the target is less than 0, no more valid
		// combinations can be created by adding to the current combination.
		if target < 0 {
			return
		}
		// Starting from 'start', explore all combinations after adding 'candidates[i]'.
		for i := start; i < len(candidates); i++ {
			// Add the current number to create a new combination.
			comb = append(comb, candidates[i])
			// Recursively explore all paths that branch from this new combination.
			backtrack(comb, i, target-candidates[i])
			// Backtrack by removing the number we just added.
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(make([]int, 0, len(candidates)), 0, target)

	return res
}
