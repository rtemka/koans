package validpalindrome

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "with digits",
			args: args{
				s: "A man, a 1plan, a canal: P1anama",
			},
			want: true,
		},
		{
			name: "with not alpha",
			args: args{
				s: "A man, a #plan, a canal: P#anama",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.args.s)
			if got != tt.want {
				t.Fatalf("validPalindrome() = %t, %t", got, tt.want)
			}
		})
	}
}
