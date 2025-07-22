// #110
// https://leetcode.com/problems/balanced-binary-tree/description/
package leetcode

import (
	btree "koans/binaryTree"
)

func isBalanced(root *btree.TreeNode) bool {
	return heightImbalance(root) != -1
}

func heightImbalance(node *btree.TreeNode) int {
	// Base case: if the node is null, it's height is 0.
	if node == nil {
		return 0
	}
	// Recursively get the height of the left and right subtrees.
	// If either subtree is imbalanced, propagate -1 up the tree.
	leftHeight := heightImbalance(node.Left)
	rightHeight := heightImbalance(node.Right)
	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}
	// If the current node's subtree is imbalanced
	// (height difference > 1), return -1.
	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	// Return the height of the current subtree.
	return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
