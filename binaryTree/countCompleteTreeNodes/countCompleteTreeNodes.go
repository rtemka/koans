// #222
// https://leetcode.com/problems/count-complete-tree-nodes/description/
package leetcode

import binarytree "koans/binaryTree"

func countNodes(root *binarytree.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
