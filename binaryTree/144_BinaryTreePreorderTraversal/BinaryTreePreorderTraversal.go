package leetcode

// Given the root of a binary tree, return the preorder traversal of its nodes' values.

// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,2,3]

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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	list := append([]int{root.Val}, preorderTraversal(root.Left)...)
	list = append(list, preorderTraversal(root.Right)...)
	return list
}

// императивно
func preorderTraversalImp(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var list []int
	var stack = []*TreeNode{root}

	for len(stack) > 0 {
		top := stack[len(stack)-1] // снимаем верхний элемент со стека
		stack = stack[:len(stack)-1]

		// добавляем текущий узел в лист
		list = append(list, top.Val)

		// добавляем правую ветвь на стэк
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		// добавляем левую ветвь на стэк
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return list
}
