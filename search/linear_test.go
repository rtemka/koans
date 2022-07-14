package search

import "testing"

func TestLinearSearchIndex(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{arr: []int{}, x: 9},
			want: -1,
		},
		{
			name: "array_have_element",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x: 9},
			want: 8,
		},
		{
			name: "array_dont_have_element",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x: 11},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearchIndex(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("LinearSearchIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
