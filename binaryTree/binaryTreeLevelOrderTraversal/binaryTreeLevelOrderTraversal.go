// #102
// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
package leetcode

import bst "koans/binaryTree"

func levelOrder(root *bst.TreeNode) [][]int {
	var levels [][]int
	return levelOrderRec(root, 0, levels)
}

func levelOrderRec(root *bst.TreeNode, level int, levels [][]int) [][]int {
	if root == nil {
		return levels
	}
	if level < len(levels) {
		levels[level] = append(levels[level], root.Val)
	} else {
		levels = append(levels, []int{root.Val})
	}
	levels = levelOrderRec(root.Left, level+1, levels)
	levels = levelOrderRec(root.Right, level+1, levels)
	return levels
}
