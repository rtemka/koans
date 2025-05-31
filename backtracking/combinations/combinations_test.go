package leetcode

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
			},
		},
		{
			name: "2",
			args: args{
				n: 2,
				k: 1,
			},
			want: [][]int{
				{1}, {2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine2(tt.args.n, tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
