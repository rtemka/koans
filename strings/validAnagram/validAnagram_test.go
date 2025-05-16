package vallidanagram

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s, t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "anagram", t: "nagaram"},
			want: true,
		},
		{
			name: "2",
			args: args{s: "rat", t: "car"},
			want: false,
		},
		{
			name: "3",
			args: args{s: "zzzap", t: "pazzz"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Fatalf("isAnagram() = %t, want %t", got, tt.want)
			}
		})
	}
}
func Test_isAnagramUnicode(t *testing.T) {
	type args struct {
		s, t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{s: "anagram", t: "nagaram"},
			want: true,
		},
		{
			name: "2",
			args: args{s: "rat", t: "car"},
			want: false,
		},
		{
			name: "3",
			args: args{s: "zzzap", t: "pazzz"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramUnicode(tt.args.s, tt.args.t); got != tt.want {
				t.Fatalf("isAnagram() = %t, want %t", got, tt.want)
			}
		})
	}
}
