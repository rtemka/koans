package leetcode

// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,3,2]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]

// Constraints:

// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

// Follow up: Recursive solution is trivial, could you do it iteratively?

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// методом рекурсии
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var list []int
	list = append(list, inorderTraversal(root.Left)...)
	list = append(list, root.Val)
	list = append(list, inorderTraversal(root.Right)...)
	return list
}

// императивно
func inorderTraversalImp(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var list []int
	var stack []*TreeNode
	var cur = root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]    // верхний элемент стека
		stack = stack[:len(stack)-1] // убираем верхний элемент

		// добавляем текущий узел в лист
		list = append(list, cur.Val)

		cur = cur.Right // идем в правую ветвь
	}

	return list
}
