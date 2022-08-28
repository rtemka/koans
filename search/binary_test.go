package search

import (
	"testing"
)

func TestBinaryIndex(t *testing.T) {
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
			if got := BinaryIndex(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("BinaryIndex() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBinary(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("findBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHas(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{arr: []int{}, x: 9},
			want: false,
		},
		{
			name: "array_have_element",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x: 9},
			want: true,
		},
		{
			name: "array_dont_have_element",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x: 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryHas(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("BinaryHas() = %v, want %v", got, tt.want)
			}
		})
	}
}
