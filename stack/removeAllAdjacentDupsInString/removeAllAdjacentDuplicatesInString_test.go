// #1047
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/
package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
