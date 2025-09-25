package leetcode

import (
	ll "koans/linkedLists"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		head  *ll.ListNode
		left  int
		right int
		want  *ll.ListNode
	}{
		{
			name:  "1",
			head:  &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 5}}}}},
			left:  2,
			right: 4,
			want:  &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 4, Next: &ll.ListNode{Val: 3, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 5}}}}},
		},
		{
			name:  "2",
			head:  &ll.ListNode{Val: 5},
			left:  1,
			right: 1,
			want:  &ll.ListNode{Val: 5},
		},
		{
			name:  "3",
			head:  &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3}}},
			left:  1,
			right: 2,
			want:  &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.head, tt.left, tt.right)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
