// #77
// https://leetcode.com/problems/combinations/description/
package leetcode

import (
	"slices"
)

func combine(n int, k int) [][]int {
	var (
		res       [][]int
		backtrack func([]int, int)
	)

	backtrack = func(comb []int, i int) {
		if len(comb) == cap(comb) {
			res = append(res, slices.Clone(comb))
			return
		}
		for i := i; i <= n; i++ {
			comb = append(comb, i)
			backtrack(comb, i+1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(make([]int, 0, k), 1)
	return res
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

// WRON ANSWER
func combine2(n int, k int) [][]int {
	var result [][]int
	for i := range n {
		for j := i; j+k <= n; j++ {
			comb := make([]int, k)

			comb[0] = nums[i]
			copy(comb[1:], nums[j+1:j+k])

			result = append(result, comb)

			if len(nums[j+1:j+k]) == 0 {
				break
			}
		}
	}
	return result
}
