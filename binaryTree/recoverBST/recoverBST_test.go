package leetcode

import (
	bst "koans/binaryTree"
	"reflect"
	"testing"
)

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		name string
		root *bst.TreeNode
		want *bst.TreeNode
	}{
		{
			name: "1",
			root: &bst.TreeNode{
				Val:  1,
				Left: &bst.TreeNode{Val: 3, Right: &bst.TreeNode{Val: 2}},
			},
			want: &bst.TreeNode{
				Val:  3,
				Left: &bst.TreeNode{Val: 1, Right: &bst.TreeNode{Val: 2}},
			},
		},
		{
			name: "2",
			root: &bst.TreeNode{
				Val:   3,
				Left:  &bst.TreeNode{Val: 1},
				Right: &bst.TreeNode{Val: 4, Left: &bst.TreeNode{Val: 2}},
			},
			want: &bst.TreeNode{
				Val:   2,
				Left:  &bst.TreeNode{Val: 1},
				Right: &bst.TreeNode{Val: 4, Left: &bst.TreeNode{Val: 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.root)
			if !reflect.DeepEqual(tt.root, tt.want) {
				t.Errorf("recoverTree() = %+v, want %+v", tt.root, tt.want)
			}
		})
	}
}
