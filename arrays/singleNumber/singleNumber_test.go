package leetcode

import "testing"

func Test_singleNumberNaive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "few",
			args: args{[]int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "few2",
			args: args{[]int{2, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberNaive(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "few",
			args: args{[]int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "few2",
			args: args{[]int{2, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
