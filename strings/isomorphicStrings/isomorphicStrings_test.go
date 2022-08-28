package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "egg_add",
			args: args{"egg", "add"},
			want: true,
		},
		{
			name: "foo_bar",
			args: args{"foo", "bar"},
			want: false,
		},
		{
			name: "paper_title",
			args: args{"paper", "title"},
			want: true,
		},
		{
			name: "badc_baba",
			args: args{"badc", "baba"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
