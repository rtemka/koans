// #283
// https://leetcode.com/problems/move-zeroes/description
package leetcode

func moveZeroes(nums []int) {
	lo := 0
	for hi := range nums {
		if nums[hi] != 0 {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			lo++
		}
	}
}
