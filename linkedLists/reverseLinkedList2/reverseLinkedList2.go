// #92
// https://leetcode.com/problems/reverse-linked-list-ii/
package leetcode

import (
	ll "koans/linkedLists"
)

func reverseBetween(head *ll.ListNode, left int, right int) *ll.ListNode {
	var (
		prev = head
		cur  = head
		i    = 1
	)
	for i < left {
		prev, cur = cur, cur.Next
		i++
	}
	lo, hi := prev, cur
	prev = nil
	for i <= right {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
		i++
	}
	switch {
	case cur != nil && left == 1: // if the reverse range is 1..i < len(list).
		lo.Next = cur
		head = prev
	case cur != nil: // if the reverse range is inbetween list.
		lo.Next = prev
		hi.Next = cur
	case left == 1: // if it is the whole list reversal.
		head = prev
	default: // if the reverse range is i..len(list).
		lo.Next = prev
	}

	return head
}
