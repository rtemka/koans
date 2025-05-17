// #27
// https://leetcode.com/problems/remove-element/description/
package leetcode

func removeElement(nums []int, val int) int {
	n := 0
	for i := range nums {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}

func removeElementNaive(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			k--
			i--
		}
	}
	return k
}
