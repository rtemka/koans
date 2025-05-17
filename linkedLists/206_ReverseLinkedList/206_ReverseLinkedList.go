// #206
// https://leetcode.com/problems/reverse-linked-list/description/
package leetcode

import ll "koans/linkedLists"

func reverseList(head *ll.ListNode) *ll.ListNode {
	var prev *ll.ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// рекурсивно
func reverseListRec(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	tail := reverseListRec(next) // доходим до конца списка
	if tail.Next == nil {        // если это последний элемент списка
		tail.Next = head // меняем указатель на текущий элемент
	} else {
		next.Next = head // иначе
	}
	return tail
}
