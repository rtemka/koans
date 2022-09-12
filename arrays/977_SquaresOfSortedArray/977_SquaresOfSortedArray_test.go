package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[-4,-1,0,3,10]=[0,1,9,16,100]",
			args: args{[]int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "[-7,-6,-5,-4,0,1,2,3]=[0,1,4,9,16,25,36,49]",
			args: args{[]int{-7, -6, -5, -4, 0, 1, 2, 3}},
			want: []int{0, 1, 4, 9, 16, 25, 36, 49},
		},
		{
			name: "[-7,-3,2,3,11]=[4,9,9,49,121]",
			args: args{[]int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[-4,-1,0,3,10]=[0,1,9,16,100]",
			args: args{[]int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "[-7,-6,-5,-4,0,1,2,3]=[0,1,4,9,16,25,36,49]",
			args: args{[]int{-7, -6, -5, -4, 0, 1, 2, 3}},
			want: []int{0, 1, 4, 9, 16, 25, 36, 49},
		},
		{
			name: "[-7,-3,2,3,11]=[4,9,9,49,121]",
			args: args{[]int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "[-7,-6,-5,-4,-3,-2,-1]=[1,4,9,16,25,36,49]",
			args: args{[]int{-7, -6, -5, -4, -3, -2, -1}},
			want: []int{1, 4, 9, 16, 25, 36, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares2Ends(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[-4,-1,0,3,10]=[0,1,9,16,100]",
			args: args{[]int{-4, -1, 0, 3, 10}},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "[-7,-6,-5,-4,0,1,2,3]=[0,1,4,9,16,25,36,49]",
			args: args{[]int{-7, -6, -5, -4, 0, 1, 2, 3}},
			want: []int{0, 1, 4, 9, 16, 25, 36, 49},
		},
		{
			name: "[-7,-3,2,3,11]=[4,9,9,49,121]",
			args: args{[]int{-7, -3, 2, 3, 11}},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "[-7,-6,-5,-4,-3,-2,-1]=[1,4,9,16,25,36,49]",
			args: args{[]int{-7, -6, -5, -4, -3, -2, -1}},
			want: []int{1, 4, 9, 16, 25, 36, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares2Ends(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares2Ends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseSquare(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverse_square",
			args: args{nums: []int{-7, -6, -5, -4, -3, -2, -1}},
			want: []int{1, 4, 9, 16, 25, 36, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseSquare(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("reverseSquare() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
