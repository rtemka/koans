package leetcode

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1,0,1,1,1]=3",
			args: args{[]int{1, 1, 0, 1, 1, 1}},
			want: 3,
		},
		{
			name: "[1,0,1,1,0,1]=2",
			args: args{[]int{1, 0, 1, 1, 0, 1}},
			want: 2,
		},
		{
			name: "[1,0]=1",
			args: args{[]int{1, 0}},
			want: 1,
		},
		{
			name: "[1]=1",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "[0]=0",
			args: args{[]int{0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnesByCount(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnesByCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
