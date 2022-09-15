package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			want: []int{1, 3, 2},
		},
		{
			name: "empty",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "only_one",
			args: args{root: &TreeNode{Val: 1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversalImp(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}},
			want: []int{1, 3, 2},
		},
		{
			name: "empty",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "only_one",
			args: args{root: &TreeNode{Val: 1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalImp(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalImp() = %v, want %v", got, tt.want)
			}
		})
	}
}
