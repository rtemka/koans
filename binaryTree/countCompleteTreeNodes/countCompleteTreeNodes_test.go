package leetcode

import (
	binarytree "koans/binaryTree"
	"testing"
)

func Test_countNodes(t *testing.T) {
	tests := []struct {
		name string
		root *binarytree.TreeNode
		want int
	}{
		{
			name: "1",
			root: &binarytree.TreeNode{
				Val:   1,
				Left:  &binarytree.TreeNode{Val: 2, Left: &binarytree.TreeNode{Val: 4}, Right: &binarytree.TreeNode{Val: 5}},
				Right: &binarytree.TreeNode{Val: 3, Left: &binarytree.TreeNode{Val: 6}},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNodes(tt.root)
			if got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
