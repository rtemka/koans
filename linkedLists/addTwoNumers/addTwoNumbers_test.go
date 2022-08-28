package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "342+465==[7,0,8]",
			args: args{l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name: "0+0==[0]",
			args: args{l1: &ListNode{Val: 0}, l2: &ListNode{Val: 0}},
			want: &ListNode{Val: 0},
		},
	}
	for _, tt := range tests {
		// res := addTwoNumbers(tt.args.l1, tt.args.l2)
		// for n := res; n != nil; n = n.Next {
		// 	t.Log(n.Val)
		// }
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbersV2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "342+465==[7,0,8]",
			args: args{l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name: "9999999+9999==[8,9,9,9,0,0,0,1]",
			args: args{l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}},
			want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}}},
		},
		{
			name: "0+0==[0]",
			args: args{l1: &ListNode{Val: 0}, l2: &ListNode{Val: 0}},
			want: &ListNode{Val: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersV2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
