// #700
// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
package leetcode

import binarytree "koans/binaryTree"

func searchBST(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func searchBSTIter(root *binarytree.TreeNode, val int) *binarytree.TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
