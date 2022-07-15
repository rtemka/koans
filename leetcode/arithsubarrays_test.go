package leetcode

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{[]int{}},
			want: 0,
		},
		{
			name: "too_small",
			args: args{[]int{1}},
			want: 0,
		},
		{
			name: "one",
			args: args{[]int{1, 2, 3}},
			want: 1,
		},
		{
			name: "three",
			args: args{[]int{1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "negative",
			args: args{[]int{-1, -5, -9, -13}},
			want: 3,
		},
		{
			name: "pos_to_neg",
			args: args{[]int{3, -1, -5, -9}},
			want: 3,
		},
		{
			name: "neg_to_pos",
			args: args{[]int{-2, 3, 8, 13, 18}},
			want: 6,
		},
		{
			name: "neg_to_pos_in_middle",
			args: args{[]int{-4, -2, 0, 2, 4, 6}},
			want: 10,
		},
		{
			name: "wrong",
			args: args{[]int{-4, -2, 3, 6, 7, 9}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
