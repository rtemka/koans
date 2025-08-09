package leetcode

import (
	bt "koans/binaryTree"
)

type Pair struct {
	Node  *bt.TreeNode
	Index int
}

func widthOfBinaryTree(root *bt.TreeNode) int {
	if root == nil {
		return 0
	}
	maxWidth := 0
	queue := []Pair{{Node: root, Index: 0}}
	for len(queue) > 0 {
		levelSize := len(queue)
		// Set the leftmost index to the index of the first node in this level. 
		// Start rightmost index at the same point as leftmost index
		// and update it as we traverse the level, eventually positioning it at the last node.
		leftmostIndex := queue[0].Index
		rightmostIndex := leftmostIndex
		for range levelSize {
			node, i := queue[0].Node, queue[0].Index
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, Pair{Node: node.Left, Index: 2*i + 1})
			}
			if node.Right != nil {
				queue = append(queue, Pair{Node: node.Right, Index: 2*i + 2})
			}
			rightmostIndex = i
		}
		maxWidth = max(maxWidth, rightmostIndex-leftmostIndex+1)
	}

	return maxWidth
}
