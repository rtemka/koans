package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123==321",
			args: args{123},
			want: 321,
		},
		{
			name: "-123==-321",
			args: args{-123},
			want: -321,
		},
		{
			name: "120==21",
			args: args{120},
			want: 21,
		},
		{
			name: "maxInt32+1==0",
			args: args{1 << 31},
			want: 0,
		},
		{
			name: "minInt32-1==0",
			args: args{-1<<31 - 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseStr(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123==321",
			args: args{123},
			want: 321,
		},
		{
			name: "-123==-321",
			args: args{-123},
			want: -321,
		},
		{
			name: "120==21",
			args: args{120},
			want: 21,
		},
		{
			name: "maxInt32+1==0",
			args: args{1 << 31},
			want: 0,
		},
		{
			name: "minInt32-1==0",
			args: args{-1<<31 - 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.x)
			if got := reverseStr(tt.args.x); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
