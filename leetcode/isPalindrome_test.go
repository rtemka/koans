package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "12222221",
			args: args{12222221},
			want: true,
		},
		{
			name: "122272221",
			args: args{122272221},
			want: true,
		},
		{
			name: "124346463",
			args: args{124346463},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
