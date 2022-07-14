package sorts

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "reversed",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "shuffled",
			args: args{[]int{2, 6, 4, 7, 1, 35, 0}},
			want: []int{0, 1, 2, 4, 6, 7, 35},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSortGenereic(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "reversed",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "shuffled",
			args: args{[]int{2, 6, 4, 7, 1, 35, 0}},
			want: []int{0, 1, 2, 4, 6, 7, 35},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSortGeneric(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
