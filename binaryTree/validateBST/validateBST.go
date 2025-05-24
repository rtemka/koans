// #98
// https://leetcode.com/problems/validate-binary-search-tree/description/
package leetcode

import (
	bst "koans/binaryTree"
	"math"
)

func isValidBST(root *bst.TreeNode) bool {
	return isValidBSTRec(root, math.MinInt, math.MaxInt)
}

func isValidBSTRec(root *bst.TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}
	return root.Val > minVal &&
		root.Val < maxVal &&
		isValidBSTRec(root.Left, minVal, root.Val) &&
		isValidBSTRec(root.Right, root.Val, maxVal)
}
