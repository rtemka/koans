package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "1",
			matrix: [][]int{
				{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
