// #128
// https://leetcode.com/problems/longest-consecutive-sequence/description/
package leetcode

func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for i := range nums {
		set[nums[i]] = true
	}
	longest := 0
	for num := range set {
		// If the current number is the smallest number in its sequence,
		// search for the length of its chain.
		if !set[num-1] {
			curNum := num
			curSeq := 1
			// Continue to find the next consecutive numbers in the sequence.
			for set[curNum+1] {
				curNum++
				curSeq++
			}
			// Update the maximum sequence so far.
			longest = max(longest, curSeq)
		}
	}
	return longest
}
