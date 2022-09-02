package leetcode

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "example3",
			args: args{strs: []string{"z"}},
			want: [][]string{{"z"}},
		},
		{
			name: "example4",
			args: args{strs: []string{"aab", "baa", "aba", "abba", "aa", "ba"}},
			want: [][]string{{"aab", "baa", "aba"}, {"abba"}, {"aa"}, {"ba"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagramsMap(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "example3",
			args: args{strs: []string{"z"}},
			want: [][]string{{"z"}},
		},
		{
			name: "example4",
			args: args{strs: []string{"aab", "baa", "aba", "abba", "aa", "ba"}},
			want: [][]string{{"aab", "baa", "aba"}, {"abba"}, {"aa"}, {"ba"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramsMap(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagramsFast(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "example2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "example3",
			args: args{strs: []string{"z"}},
			want: [][]string{{"z"}},
		},
		{
			name: "example4",
			args: args{strs: []string{"aab", "baa", "aba", "abba", "aa", "ba"}},
			want: [][]string{{"aab", "baa", "aba"}, {"abba"}, {"aa"}, {"ba"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramsFast(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramsFast() = %v, want %v", got, tt.want)
			}
		})
	}
}
