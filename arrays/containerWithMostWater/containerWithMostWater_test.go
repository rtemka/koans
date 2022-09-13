package leetcode

import "testing"

func Test_maxAreaSlow(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,8,6,2,5,4,8,3,7]=49",
			args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "[1,2,1]=3",
			args: args{[]int{1, 2, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaSlow(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,8,6,2,5,4,8,3,7]=49",
			args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "[1,2,1]=3",
			args: args{[]int{1, 2, 1}},
			want: 2,
		},
		{
			name: "[3,9,3,4,7,2,12,6]=45",
			args: args{[]int{3, 9, 3, 4, 7, 2, 12, 6}},
			want: 45,
		},
		{
			name: "[3,90,90,4,7,2,12,6]=90",
			args: args{[]int{3, 90, 90, 4, 7, 2, 12, 6}},
			want: 90,
		},
		{
			name: "[2,1]=1",
			args: args{[]int{2, 1}},
			want: 1,
		},
		{
			name: "[1,2,3,4,5,25,24,3,4]=24",
			args: args{[]int{1, 2, 3, 4, 5, 25, 24, 3, 4}},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
