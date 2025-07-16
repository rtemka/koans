// #496
// https://leetcode.com/problems/next-greater-element-i/description/
package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// Save the indexes of nums1 in the set to lookup.
	set := make(map[int]int, len(nums1))
	for i := range nums1 {
		set[nums1[i]] = i
	}
	var stack []int
	// Find the next largest number of each element, starting with the rightmost element.
	for i := len(nums2) - 1; i >= 0; i-- {
		// Pop values from the top of the stack until the current
		// value's next largest number is at the top.
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		// If set has candidate number record the current value's next largest number,
		// which is at the top of the stack. If the stack is empty, record -1.
		if k, ok := set[nums2[i]]; ok {
			if len(stack) > 0 {
				nums1[k] = stack[len(stack)-1]
			} else {
				nums1[k] = -1
			}
		}
		// Append current value to the top of the stack.
		stack = append(stack, nums2[i])
	}
	return nums1
}
