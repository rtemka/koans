package sorts

import "testing"

func Test_majorityElementNaive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{[]int{3, 2, 3}},
			want: 3,
		},
		{
			name: "second",
			args: args{[]int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			name: "third",
			args: args{[]int{6, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementNaive(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{[]int{3, 2, 3}},
			want: 3,
		},
		{
			name: "second",
			args: args{[]int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			name: "third",
			args: args{[]int{6, 5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
