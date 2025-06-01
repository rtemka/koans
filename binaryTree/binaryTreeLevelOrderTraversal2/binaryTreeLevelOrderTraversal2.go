// #107
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
package leetcode

import (
	bst "koans/binaryTree"
	"slices"
)

func levelOrderBottom(root *bst.TreeNode) [][]int {
	levels := levelOrderRec(root, 0, [][]int{})
	slices.Reverse(levels)
	return levels
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
