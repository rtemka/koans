// #239
// https://leetcode.com/problems/sliding-window-maximum/description/
package leetcode

import "container/list"

type Tuple struct {
	Val int
	Idx int
}

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums))
	lo, hi, n, dq := 0, 0, len(nums), list.New()
	for hi < n {
		// Ensure the values of the deque maintain a monotonic decreasing order
		// by removing candidates ≤ the current candidate.
		for dq.Len() > 0 && dq.Back().Value.(Tuple).Val <= nums[hi] {
			dq.Remove(dq.Back())
		}
		// Add the current candidate.
		dq.PushBack(Tuple{Val: nums[hi], Idx: hi})
		// If the window is of length 'k', record the maximum of the window.
		if hi-lo+1 == k {
			// Remove values whose indexes occur outside the window.
			if el := dq.Front(); el != nil && el.Value.(Tuple).Idx < lo {
				dq.Remove(el)
			}
			// The maximum value of this window is the front value in the deque.
			result = append(result, dq.Front().Value.(Tuple).Val)
			// Slide the window by advancing both 'lo' and 'hi'.
			// The `hi` pointer always gets advanced so we just need to advance 'lo'.
			lo++
		}
		hi++
	}
	return result
}
