package leetcode

import (
	bst "koans/binaryTree"
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
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
			want: [][]int{{15, 7}, {9, 20}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrderBottom(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
