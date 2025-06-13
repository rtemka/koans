package leetcode

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{
			name:    "1",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0,
			dst:     3,
			k:       1,
			want:    700,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			if got != tt.want {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
