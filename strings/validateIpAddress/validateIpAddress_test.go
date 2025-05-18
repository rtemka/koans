package leetcode

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		queryIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				queryIP: "172.16.254.1",
			},
			want: "IPv4",
		},
		{
			name: "2",
			args: args{
				queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334",
			},
			want: "IPv6",
		},
		{
			name: "3",
			args: args{
				queryIP: "256.256.256.256",
			},
			want: "Neither",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validIPAddress(tt.args.queryIP)
			if got != tt.want {
				t.Errorf("validIPAddress() = %s, want %s", got, tt.want)
			}
		})
	}
}
