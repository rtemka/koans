package leetcode

import "testing"

func Test_networkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n     int
		k     int
		want  int
	}{
		{
			name: "1",
			times: [][]int{
				{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
			},
			n:    4,
			k:    2,
			want: 2,
		},
		{
			name: "2",
			times: [][]int{
				{1, 2, 1},
			},
			n:    2,
			k:    1,
			want: 1,
		},
		{
			name: "3",
			times: [][]int{
				{1, 2, 1},
			},
			n:    2,
			k:    2,
			want: -1,
		},
		{
			name: "4",
			times: [][]int{
				{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59},
				{5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43},
				{4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16},
				{3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0},
				{2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59},
			},
			n:    5,
			k:    5,
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := networkDelayTime(tt.times, tt.n, tt.k)
			if got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
