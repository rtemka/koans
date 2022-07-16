package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{"()"},
			want: true,
		},
		{
			name: "invalid",
			args: args{"("},
			want: false,
		},
		{
			name: "valid_long",
			args: args{"()[]{}"},
			want: true,
		},
		{
			name: "invalid_long",
			args: args{"()[]{})"},
			want: false,
		},
		{
			name: "invalid_long_mix",
			args: args{"()[]}{"},
			want: false,
		},
		{
			name: "valid_long_mix",
			args: args{"([{}])"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
