// #230
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
package leetcode

import bt "koans/binaryTree"

func kthSmallest(root *bt.TreeNode, k int) int {
	var (
		arr     []int
		inorder func(root *bt.TreeNode)
	)

	inorder = func(root *bt.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		arr = append(arr, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return arr[k-1]
}
