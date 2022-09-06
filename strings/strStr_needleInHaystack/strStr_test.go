package leetcode

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hello<<ll=2",
			args: args{haystack: "hello", needle: "ll"},
			want: 2,
		},
		{
			name: "empty_needle",
			args: args{haystack: "hello", needle: ""},
			want: 0,
		},
		{
			name: "aaaaa<<bba=-1",
			args: args{haystack: "aaaaa", needle: "bba"},
			want: -1,
		},
		{
			name: "aaabba<<bba=",
			args: args{haystack: "aaabba", needle: "bba"},
			want: 3,
		},
		{
			name: "aaabbz<<z=",
			args: args{haystack: "aaabbz", needle: "z"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStrCheating(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hello<<ll=2",
			args: args{haystack: "hello", needle: "ll"},
			want: 2,
		},
		{
			name: "empty_needle",
			args: args{haystack: "hello", needle: ""},
			want: 0,
		},
		{
			name: "aaaaa<<bba=-1",
			args: args{haystack: "aaaaa", needle: "bba"},
			want: -1,
		},
		{
			name: "aaabba<<bba=",
			args: args{haystack: "aaabba", needle: "bba"},
			want: 3,
		},
		{
			name: "aaabbz<<z=",
			args: args{haystack: "aaabbz", needle: "z"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrCheating(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStrCheating() = %v, want %v", got, tt.want)
			}
		})
	}
}
