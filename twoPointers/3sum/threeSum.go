// #15
// https://leetcode.com/problems/3sum/description/
package leetcode

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)

	for i := range nums {
		// Triplets consisting of only positive numbers will never sum to 0.
		if nums[i] > 0 {
			break
		}
		// To avoid duplicate triplets, skip first triplet element
		// if it's the same as the previous number.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Find all pairs that sum to a -nums[i].
		target := -nums[i]
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			if sum == target {
				result = append(result, []int{nums[i], nums[lo], nums[hi]})
				lo++
				// To avoid duplicate '[nums[lo], nums[hi]]' pairs,
				// skip nums[lo] if it's the same as the previous number.
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			} else if sum < target {
				// The left pointer will always point to a value less than or equal
				// to the value at the right pointer because the array is sorted.
				// Incrementing it would result in a sum greater than or equal to the current sum.
				// So we aiming to increase the sum toward the target value.
				lo++
			} else {
				// Decrementing the right pointer would result in a sum that's less than or equal to current sum.
				// So we aiming to decrease the sum toward the target value.
				hi--
			}
		}
	}
	return result
}
