// #226
// https://leetcode.com/problems/invert-binary-tree/description/
package leetcode

import bst "koans/binaryTree"

func invertTree(root *bst.TreeNode) *bst.TreeNode {
	// Base case: If the node is null, there's nothing to invert.
	if root == nil {
		return nil
	}
	// Swap the left and right subtrees of the current node.
	root.Left, root.Right = root.Right, root.Left
	// Recursively invert the left and right subtrees.
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTreeIterative(root *bst.TreeNode) *bst.TreeNode {
	if root == nil {
		return nil
	}
	stack := []*bst.TreeNode{root}
	for len(stack) > 0 {
		// Pop of the stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Swap the left and right subtrees of the current node.
		node.Left, node.Right = node.Right, node.Left
		// Push the left and right subtrees onto the stack.
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}
