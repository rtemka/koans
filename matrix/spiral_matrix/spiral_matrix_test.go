package leetcode

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			},
			want: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			name: "4",
			args: args{
				matrix: [][]int{{6, 9, 7}},
			},
			want: []int{6, 9, 7},
		},
		{
			name: "5",
			args: args{
				matrix: [][]int{{6}, {9}, {7}},
			},
			want: []int{6, 9, 7},
		},
		{
			name: "6",
			args: args{
				matrix: [][]int{{1}},
			},
			want: []int{1},
		},
		{
			name: "7",
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
			},
			want: []int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		got := spiralOrder(tt.args.matrix)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
		}
	}
}
