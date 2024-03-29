package leetcode

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ab>>ba",
			args: args{"ab", "ba"},
			want: true,
		},
		{
			name: "ab>>ab",
			args: args{"ab", "ab"},
			want: false,
		},
		{
			name: "aa>>aa",
			args: args{"aa", "aa"},
			want: true,
		},
		{
			name: "ragep>>pager",
			args: args{"ragep", "pager"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
