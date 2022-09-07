package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,1,2,4,3]=2",
			args: args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}},
			want: 2,
		},
		{
			name: "[1,4,4]=1",
			args: args{target: 4, nums: []int{1, 4, 4}},
			want: 1,
		},
		{
			name: "[10,5,13,4,8,4,5,11,14,9,16,10,20,8]=6",
			args: args{target: 80, nums: []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}