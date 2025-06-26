// #160
// https://leetcode.com/problems/intersection-of-two-linked-lists/description/
package leetcode

import (
	ll "koans/linkedLists"
)

func getIntersectionNode(headA, headB *ll.ListNode) *ll.ListNode {
	a, b := headA, headB
	// Traverse through the LinkedList 'a' and 'b' until they meet.
	for a != b {
		// Traverse whole list 'a' and then at the end of list 'a' start to traverse list 'b'.
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		// At the same time
		// Traverse whole list 'b' and then at the end of list 'b' start to traverse list 'a'.
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	// Here 'a' and 'b' either both nil(lists don't intersect) or at the intersection node.
	return a
}
