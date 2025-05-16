package leetcode

import (
	binarytree "koans/binaryTree"
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *binarytree.TreeNode
	}{
		{
			name: "1",
			args: args{
				root: &binarytree.TreeNode{
					Val: 4,
					Left: &binarytree.TreeNode{
						Val:   2,
						Left:  &binarytree.TreeNode{Val: 1},
						Right: &binarytree.TreeNode{Val: 3},
					},
					Right: &binarytree.TreeNode{Val: 7},
				},
				val: 2,
			},
			want: &binarytree.TreeNode{
				Val:   2,
				Left:  &binarytree.TreeNode{Val: 1},
				Right: &binarytree.TreeNode{Val: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
