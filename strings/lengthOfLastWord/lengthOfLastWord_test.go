package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{""},
			want: 0,
		},
		{
			name: "hello_world",
			args: args{"Hello World"},
			want: 5,
		},
		{
			name: "to_the_moon",
			args: args{"   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "joyboy",
			args: args{"luffy is still joyboy"},
			want: 6,
		},
		{
			name: "one_word",
			args: args{"luffy"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
