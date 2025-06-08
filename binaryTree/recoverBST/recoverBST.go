// #99
// https://leetcode.com/problems/recover-binary-search-tree/description/
package leetcode

import (
	bst "koans/binaryTree"
	"slices"
)

func recoverTree(root *bst.TreeNode) {
	var first, second, prev *bst.TreeNode
	var inOrder func(*bst.TreeNode)

	inOrder = func(root *bst.TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		// Если предыдущий узел не пустой и значение текущего узла меньше значения предыдущего узла,
		// то мы нашли два узла, которые нарушают порядок возрастания.
		if prev != nil {
			// Если первый узел еще не был найден, то мы его запоминаем.
			if first == nil && prev.Val >= root.Val {
				first = prev
			}
			// Запоминаем второй узел, когда первый уже найден.
			if first != nil && prev.Val >= root.Val {
				second = root
			}
		}
		// После того, как мы обработали текущий узел, мы запоминаем его как предыдущий.
		prev = root
		// Обрабатываем правое поддерево.
		inOrder(root.Right)
	}

	inOrder(root)

	first.Val, second.Val = second.Val, first.Val
}

func RecoverTreeNlogN(root *bst.TreeNode) {
	// Get all the numbers.
	var vals []int
	inOrder(root, func(node *bst.TreeNode) {
		vals = append(vals, node.Val)
	})
	// Sort them.
	slices.Sort(vals)
	i := 0
	// Put them in order.
	inOrder(root, func(node *bst.TreeNode) {
		node.Val = vals[i]
		i++
	})
}
func inOrder(root *bst.TreeNode, f func(node *bst.TreeNode)) {
	if root == nil {
		return
	}
	inOrder(root.Left, f)
	f(root)
	inOrder(root.Right, f)
}
