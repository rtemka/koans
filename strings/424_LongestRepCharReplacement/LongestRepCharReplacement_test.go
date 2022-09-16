package leetcode

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ABAB;k=2==4",
			args: args{"ABAB", 2},
			want: 4,
		},
		{
			name: "AABABBA;k=1==4",
			args: args{"AABABBA", 1},
			want: 4,
		},
		{
			name: "empty;k=4==0",
			args: args{"", 4},
			want: 0,
		},
		{
			name: "ABBB;k=2==4",
			args: args{"ABBB", 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
