package leetcode

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,5]=1",
			args: args{[]int{1, 3, 5}},
			want: 1,
		},
		{
			name: "[2,2,2,0,1]=0",
			args: args{[]int{2, 2, 2, 0, 1}},
			want: 0,
		},
		{
			name: "[3,3,3,3,1,3]=1",
			args: args{[]int{3, 3, 3, 3, 1, 3}},
			want: 1,
		},
		{
			name: "[1,3,3]=1",
			args: args{[]int{1, 3, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
