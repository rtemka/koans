package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {

	n := &ListNode{Val: 2}
	cycleList := &ListNode{Val: 1, Next: n}
	cycleList.Next.Next = &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: n}}

	noCycleList := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_cycle",
			args: args{cycleList},
			want: true,
		},
		{
			name: "no_cycle",
			args: args{noCycleList},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasCycle2(t *testing.T) {

	n := &ListNode{Val: 2}
	cycleList := &ListNode{Val: 1, Next: n}
	cycleList.Next.Next = &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: n}}

	noCycleList := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has_cycle",
			args: args{cycleList},
			want: true,
		},
		{
			name: "no_cycle",
			args: args{noCycleList},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle2(tt.args.head); got != tt.want {
				t.Errorf("hasCycle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
