package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{{1, 8}},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "4",
			args: args{
				intervals: [][]int{{1, 8}, {2, 4}, {3, 6}, {6, 7}},
			},
			want: [][]int{{1, 8}},
		},
		{
			name: "5",
			args: args{
				intervals: [][]int{{1, 4}, {5, 7}, {8, 9}},
			},
			want: [][]int{{1, 4}, {5, 7}, {8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.intervals)
			t.Log(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
