package leetcode

import (
	bt "koans/binaryTree"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *bt.TreeNode
		k    int
		want int
	}{
		{
			name: "1",
			root: &bt.TreeNode{
				Val: 5,
				Left: &bt.TreeNode{
					Val:   2,
					Left:  &bt.TreeNode{Val: 1},
					Right: &bt.TreeNode{Val: 4},
				},
				Right: &bt.TreeNode{
					Val:   7,
					Left:  &bt.TreeNode{Val: 6},
					Right: &bt.TreeNode{Val: 9},
				},
			},
			k:    5,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallest(tt.root, tt.k)
			if got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
