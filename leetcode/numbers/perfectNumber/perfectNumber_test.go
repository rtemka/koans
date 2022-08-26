package leetcode

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "28==true",
			args: args{num: 28},
			want: true,
		},
		{
			name: "7==false",
			args: args{num: 7},
			want: false,
		},
		{
			name: "496==true",
			args: args{num: 496},
			want: true,
		},
		{
			name: "315==false",
			args: args{num: 315},
			want: false,
		},
		{
			name: "99999997==false",
			args: args{num: 99999997},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
