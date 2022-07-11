package strings

import "testing"

func TestCountRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero string",
			args: args{""},
			want: 0,
		},
		{
			name: "regular ASCII string",
			args: args{"this is regular string"},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountRunes(tt.args.s); got != tt.want {
				t.Errorf("CountRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
