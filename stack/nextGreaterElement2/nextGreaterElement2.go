// #503
// https://leetcode.com/problems/next-greater-element-ii/description/
package leetcode

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	var stack []int
	// Find the next largest number of each element, starting with the rightmost element.
	// But this time as opposed to `next greater element 1` we need to process `nums` 2 times.
	for i := len(nums)*2 - 1; i >= 0; i-- {
		// Determine the index in the `nums` and in the `res`.
		idx := i % len(nums)
		for len(stack) > 0 && stack[len(stack)-1] <= nums[idx] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[idx] = stack[len(stack)-1]
		} else {
			res[idx] = -1
		}
		stack = append(stack, nums[idx])
	}
	return res
}
