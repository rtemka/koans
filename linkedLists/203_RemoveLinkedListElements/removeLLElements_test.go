package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	tl := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}
	wl := tl
	wl.Next.Next = tl.Next.Next.Next
	wl.Next.Next.Next.Next.Next = nil

	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[1,2,6,3,4,5,6],val=6==[1,2,3,4,5]",
			args: args{head: tl, val: 6},
			want: wl,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
