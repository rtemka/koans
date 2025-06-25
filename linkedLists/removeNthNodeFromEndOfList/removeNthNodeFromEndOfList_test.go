package leetcode

import (
	ll "koans/linkedLists"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ll.ListNode
		n    int
		want *ll.ListNode
	}{
		{
			name: "1",
			head: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5}}}}},
			n:    2,
			want: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 5}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.head, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
