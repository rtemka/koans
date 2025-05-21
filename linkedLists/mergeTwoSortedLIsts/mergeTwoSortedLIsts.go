// #21
// https://leetcode.com/problems/merge-two-sorted-lists/description/
package leetcode

import ll "koans/linkedLists"

func mergeTwoLists(list1 *ll.ListNode, list2 *ll.ListNode) *ll.ListNode {
	cur := &ll.ListNode{}
	head := cur
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	for list1 != nil {
		cur.Next = list1
		cur = cur.Next
		list1 = list1.Next
	}
	for list2 != nil {
		cur.Next = list2
		cur = cur.Next
		list2 = list2.Next
	}
	return head.Next
}
