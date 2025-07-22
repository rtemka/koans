package leetcode

import (
	btree "koans/binaryTree"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *btree.TreeNode
		want bool
	}{
		{
			name: "1",
			root: &btree.TreeNode{Val: 3, Left: &btree.TreeNode{Val: 9}, Right: &btree.TreeNode{Val: 20, Left: &btree.TreeNode{Val: 15}, Right: &btree.TreeNode{Val: 7}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tf := isBalanced(tt.root); tf != tt.want {
				t.Errorf("isBalanced() = %v, want %v", tf, tt.want)
			}
		})
	}
}
