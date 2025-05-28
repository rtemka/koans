// #47
// https://leetcode.com/problems/permutations-ii/description/
// Explanation: https://leetcode.com/problems/permutations-ii/editorial/
package leetcode

import "slices"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	memo := make(map[int]int, len(nums))
	for _, num := range nums {
		memo[num]++
	}
	var backtrack func(comb []int, size int)

	backtrack = func(comb []int, size int) {
		if len(comb) == size {
			result = append(result, slices.Clone(comb))
			return
		}
		for num, cnt := range memo {
			if cnt == 0 {
				continue
			}

			comb = append(comb, num)
			memo[num]--
			backtrack(comb, size)
			memo[num]++
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(make([]int, 0, len(nums)), len(nums))
	return result
}
