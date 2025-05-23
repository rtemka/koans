package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "2",
			args: args{
				n: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		got := generateMatrix(tt.args.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
		}
	}
}
