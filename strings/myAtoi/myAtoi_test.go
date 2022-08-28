package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "42",
			args: args{"42"},
			want: 42,
		},
		{
			name: "-42",
			args: args{"-42"},
			want: -42,
		},
		{
			name: "4193_with_words",
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: "words_and_987",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "+-42",
			args: args{"+-42"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
