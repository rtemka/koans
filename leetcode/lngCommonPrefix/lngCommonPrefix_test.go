package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "одинаковые_слова",
			args: args{[]string{"football", "football", "football"}},
			want: "football",
		},
		{
			name: "слова_разной_длины",
			args: args{[]string{"money", "honey", "football"}},
			want: "",
		},
		{
			name: "пустые",
			args: args{[]string{"", "", ""}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %q, want %q", got, tt.want)
			}
		})
	}
}
