// #101
// https://leetcode.com/problems/symmetric-tree/description
package leetcode

import (
	bt "koans/binaryTree"
)

func isSymmetric(root *bt.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRec(root.Left, root.Right)
}

func isSymmetricRec(left *bt.TreeNode, right *bt.TreeNode) bool {
	switch {
	case left == nil && right == nil:
		return true
	case left != nil && right != nil && left.Val == right.Val:
		return isSymmetricRec(left.Left, right.Right) && isSymmetricRec(left.Right, right.Left)
	default:
		return false
	}
}
