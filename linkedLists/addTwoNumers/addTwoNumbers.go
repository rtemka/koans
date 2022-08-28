package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]
// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// 9999999
// 9999
// 89991111
// 455
//  07

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return fromInt(toInt(l1) + toInt(l2))
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {

	root := &ListNode{}
	for c, n, node := 0, 0, root; l1 != nil || l2 != nil || c > 0; {

		switch {
		case l1 != nil && l2 != nil:
			n = l1.Val + l2.Val + c
			l1, l2 = l1.Next, l2.Next
		case l1 != nil:
			n = l1.Val + c
			l1 = l1.Next
		case l2 != nil:
			n = l2.Val + c
			l2 = l2.Next
		case c > 0:
			n = c
		}

		if n >= 10 {
			c = 1
			node.Val = n - 10
		} else {
			c = 0
			node.Val = n
		}
		if l1 != nil || l2 != nil || c > 0 {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	return root
}

func toInt(l *ListNode) uint64 {
	var xs []int

	for n := l; n != nil; n = n.Next {
		xs = append(xs, n.Val)
	}

	var n uint64
	for i := len(xs) - 1; i >= 0; i-- {
		if i != len(xs)-1 {
			n = n*10 + uint64(xs[i])
		} else {
			n = uint64(xs[i])
		}
	}

	return n
}

func fromInt(n uint64) *ListNode {

	root := &ListNode{}
	for i, node := n, root; i != 0; i, node = i/10, node.Next {
		node.Val = int(i % 10)
		if i/10 != 0 {
			node.Next = &ListNode{}
		}
	}

	return root
}
