package leetcode

import (
	ll "koans/linkedLists"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ll.ListNode
		list2 *ll.ListNode
	}
	tests := []struct {
		name string
		args args
		want *ll.ListNode
	}{
		{
			name: "1",
			args: args{
				list1: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 4}}},
				list2: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4}}},
			},
			want: &ll.ListNode{
				Val: 1, Next: &ll.ListNode{
					Val: 1, Next: &ll.ListNode{
						Val: 2, Next: &ll.ListNode{
							Val: 3, Next: &ll.ListNode{
								Val: 4, Next: &ll.ListNode{
									Val: 4}}}}}},
		},
	}

	for _, tt := range tests {
		got := mergeTwoLists(tt.args.list1, tt.args.list2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
		}
	}
}
