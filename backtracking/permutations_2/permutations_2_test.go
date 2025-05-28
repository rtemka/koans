// #47
// https://leetcode.com/problems/permutations-ii/description/
package leetcode

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{-10, 4},
			},
			want: [][]int{
				{-10, 4}, {4, -10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
