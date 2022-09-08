package leetcode

// Given the head of a linked list and an integer val, remove all the nodes of
// the linked list that has Node.val == val, and return the new head.

// Example 1:
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]

// Example 2:

// Input: head = [], val = 1
// Output: []

// Example 3:

// Input: head = [7,7,7,7], val = 7
// Output: []

// Constraints:

// The number of nodes in the list is in the range [0, 104].
// 1 <= Node.val <= 50
// 0 <= val <= 50

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	// определяем head связного списка
	var h = &ListNode{Next: head}

	for prev, cur := h, head; cur != nil; {
		// если текущий. не nil и его значение равно
		// искомому, то пробрасываем ссылку вокруг него
		if cur.Val == val {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return h.Next
}

func RemoveElements2(head *ListNode, val int) *ListNode {
	// определяем head связного списка
	var h = &ListNode{Val: head.Val, Next: head.Next}
	for h != nil && h.Val == val {
		h = h.Next
	}

	for n := h; n != nil; n = n.Next {
		// если след. не nil и его значение равно
		// искомому, то пробрасываем ссылку вокруг него
		if n.Next != nil && n.Val == val {
			n.Next = n.Next.Next
		}
	}

	return h
}
