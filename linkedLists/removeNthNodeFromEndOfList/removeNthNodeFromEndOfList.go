// #19
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
package leetcode

import ll "koans/linkedLists"

func removeNthFromEnd(head *ll.ListNode, n int) *ll.ListNode {
	// Make a dummy node for the case when we need to remove 'head' node.
	dummy := &ll.ListNode{Next: head}
	// Initialize the two pointers - one will be 'n' steps ahead of another.
	lo, hi := dummy, dummy
	// Constraints:
	// 1 <= n <= NumberOfNodesInList
	// We can advance the 'hi' pointer not afraiding
	// that 'n' is larger then number of nodes in the list.
	for range n {
		hi = hi.Next
	}
	// Then move both pointers toward the end of the linked list.
	for hi.Next != nil {
		hi, lo = hi.Next, lo.Next
	}
	// Finally the 'lo' pointer is on the 'n-1' position in the list.
	// Remove the node next to the 'lo' pointer.
	lo.Next = lo.Next.Next

	return dummy.Next
}
