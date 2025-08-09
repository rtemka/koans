package leetcode

import (
	bt "koans/binaryTree"
	"testing"
)

func Test_widthOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *bt.TreeNode
		want int
	}{
		{
			name: "1",
			root: &bt.TreeNode{Val: 1, Left: &bt.TreeNode{Val: 3, Left: &bt.TreeNode{Val: 5}, Right: &bt.TreeNode{Val: 3}}, Right: &bt.TreeNode{Val: 2, Right: &bt.TreeNode{Val: 9}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := widthOfBinaryTree(tt.root)
			if got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
