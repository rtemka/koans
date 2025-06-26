package leetcode

import (
	ll "koans/linkedLists"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	intersectionNode := &ll.ListNode{Val: 100, Next: &ll.ListNode{Val: 99, Next: &ll.ListNode{Val: 88}}}
	tests := []struct {
		name  string
		headA *ll.ListNode
		headB *ll.ListNode
		want  *ll.ListNode
	}{
		{
			name:  "1",
			headA: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: intersectionNode}},
			headB: &ll.ListNode{Val: 1, Next: &ll.ListNode{Val: 2, Next: &ll.ListNode{Val: 3, Next: intersectionNode}}},
			want:  intersectionNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntersectionNode(tt.headA, tt.headB)
			if got != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
