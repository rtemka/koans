package strings

import "testing"

func TestСontains(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{"", ""},
			want: false,
		},
		{
			name: "hello in 'hello world'",
			args: args{"hello world", "hello"},
			want: true,
		},
		{
			name: "world in 'hello world'",
			args: args{"hello world", "world"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Сontains(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("Сontains() = %v, want %v", got, tt.want)
			}
		})
	}
}
