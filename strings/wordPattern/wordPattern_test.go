package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abba_true",
			args: args{"abba", "dog cat cat dog"},
			want: true,
		},
		{
			name: "abba_false",
			args: args{"abba", "dog cat cat fish"},
			want: false,
		},
		{
			name: "abba_false_one_word",
			args: args{"abba", "dog dog dog dog"},
			want: false,
		},
		{
			name: "aaaa_false",
			args: args{"aaaa", "dog cat cat dog"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
