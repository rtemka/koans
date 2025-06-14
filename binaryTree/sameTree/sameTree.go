// #100
// https://leetcode.com/problems/same-tree/description
package leetcode

import bt "koans/binaryTree"

func isSameTree(p *bt.TreeNode, q *bt.TreeNode) bool {
	switch {
	case p == nil && q == nil:
		return true
	case p != nil && q != nil && p.Val == q.Val:
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	default:
		return false
	}
}
