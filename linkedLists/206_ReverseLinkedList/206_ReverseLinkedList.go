package leetcode

// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

// Example 2:
// Input: head = [1,2]
// Output: [2,1]

// Example 3:
// Input: head = []
// Output: []

// Constraints:
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000

// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// императивно
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev, cur := head, head.Next
	head.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

// рекурсивно
func reverseListRec(head *ListNode) *ListNode {
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
