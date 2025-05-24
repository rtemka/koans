package leetcode

import (
	binarytree "koans/binaryTree"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *binarytree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root: &binarytree.TreeNode{
					Val:   2,
					Left:  &binarytree.TreeNode{Val: 1},
					Right: &binarytree.TreeNode{Val: 3},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %t, want %t", got, tt.want)
			}
		})
	}
}
