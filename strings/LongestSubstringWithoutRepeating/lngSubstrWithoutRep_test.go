package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb=3",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "bbbbb=1",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "pwwkew=3",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "dvdf=3",
			args: args{"dvdf"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
