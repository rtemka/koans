package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
			want: []int{1, 2, 3},
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
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversalImp(t *testing.T) {
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
			want: []int{1, 2, 3},
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
			if got := preorderTraversalImp(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversalImp() = %v, want %v", got, tt.want)
			}
		})
	}
}
