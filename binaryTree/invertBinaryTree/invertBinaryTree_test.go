package leetcode

import (
	bst "koans/binaryTree"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string
		root *bst.TreeNode
		want *bst.TreeNode
	}{
		{
			name: "1",
			root: &bst.TreeNode{Val: 4, Left: &bst.TreeNode{Val: 2, Left: &bst.TreeNode{Val: 1}, Right: &bst.TreeNode{Val: 3}}, Right: &bst.TreeNode{Val: 7, Left: &bst.TreeNode{Val: 6}, Right: &bst.TreeNode{Val: 9}}},
			want: &bst.TreeNode{Val: 4, Left: &bst.TreeNode{Val: 7, Left: &bst.TreeNode{Val: 9}, Right: &bst.TreeNode{Val: 6}}, Right: &bst.TreeNode{Val: 2, Left: &bst.TreeNode{Val: 3}, Right: &bst.TreeNode{Val: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
