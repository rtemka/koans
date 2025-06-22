package leetcode

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "1",
			matrix: [][]int{
				{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
			},
			want: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
