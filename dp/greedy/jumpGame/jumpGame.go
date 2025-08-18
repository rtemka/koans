// #55
// https://leetcode.com/problems/jump-game/description/
package leetcode

func canJump(nums []int) bool {
	// Set the initial destination to the last index in the array.
	destination := len(nums) - 1
	// Traverse the array in reverse to see if the destination can be reached by earlier indexes.
	for i := destination; i >= 0; i-- {
		// If we can reach the destination from the current index, set this index as the new destination.
		if i+nums[i] >= destination {
			destination = i
		}
	}
	// If the destination is index 0, we can jump to the end from index 0.
	return destination == 0
}
