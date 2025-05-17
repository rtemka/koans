package leetcode

import (
	ll "koans/linkedLists"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tl := &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2}}
	type args struct {
		head *ll.ListNode
	}
	tests := []struct {
		name string
		args args
		want *ll.ListNode
	}{
		{
			name: "[1,2]=[2,1]",
			args: args{tl},
			want: tl.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			printList(t, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListRec(t *testing.T) {
	tl := &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2}}
	tl1 := &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4}}}}
	tl2 := tl1.Next.Next.Next

	type args struct {
		head *ll.ListNode
	}
	tests := []struct {
		name string
		args args
		want *ll.ListNode
	}{
		{
			name: "[1,2]=[2,1]",
			args: args{tl},
			want: tl.Next,
		},
		{
			name: "[1,2,3,4]=[4,3,2,1]",
			args: args{tl1},
			want: tl2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseListRec(tt.args.head)
			printList(t, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func printList(t *testing.T, head *ll.ListNode) {
	t.Helper()
	for n := head; n != nil; n = n.Next {
		if n.Next != nil && n.Next.Next == n {
			t.Log("cycle!")
			return
		}
		t.Log(n)
	}
}
