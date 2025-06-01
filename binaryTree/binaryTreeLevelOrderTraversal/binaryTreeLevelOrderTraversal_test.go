package leetcode

import (
	bst "koans/binaryTree"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *bst.TreeNode
		want [][]int
	}{
		{
			name: "1",
			root: &bst.TreeNode{
				Val: 3, Left: &bst.TreeNode{Val: 9}, Right: &bst.TreeNode{Val: 20, Left: &bst.TreeNode{Val: 15}, Right: &bst.TreeNode{Val: 7}},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
