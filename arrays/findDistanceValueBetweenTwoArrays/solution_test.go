package leetcode

import "testing"

func Test_findTheDistanceValue(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr1: []int{4, 5, 8},
				arr2: []int{10, 9, 1, 8},
				d:    2,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				arr1: []int{10, 12, 13, 14, 15},
				arr2: []int{10, 9, 1, 8},
				d:    2,
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				arr1: []int{-10, -12, -13, -14, -15},
				arr2: []int{10, 9, 1, 8},
				d:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue(tt.args.arr1, tt.args.arr2, tt.args.d); got != tt.want {
				t.Fatalf("findTheDistanceValue() = %d, want %d", got, tt.want)
			}
		})
	}
}
