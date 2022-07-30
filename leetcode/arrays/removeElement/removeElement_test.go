package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{nums: []int{}, val: 2},
			want: 0,
		},
		{
			name: "remove_five",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, val: 5},
			want: 9,
		},
		{
			name: "remove_three",
			args: args{nums: []int{3, 2, 2, 3}, val: 3},
			want: 2,
		},
		{
			name: "remove_two",
			args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
				t.Log(tt.args.nums)
			}
		})
	}
}

func Test_removeElementNaive(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{nums: []int{}, val: 2},
			want: 0,
		},
		{
			name: "remove_five",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, val: 5},
			want: 9,
		},
		{
			name: "remove_three",
			args: args{nums: []int{3, 2, 2, 3}, val: 3},
			want: 2,
		},
		{
			name: "remove_two",
			args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementNaive(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
				t.Log(tt.args.nums)
			}
		})
	}
}
