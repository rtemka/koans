package leetcode

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1.01==1.001",
			args: args{"1.01", "1.001"},
			want: 0,
		},
		{
			name: "1.0==1.0.0",
			args: args{"1.0", "1.0.0"},
			want: 0,
		},
		{
			name: "0.1==1.1",
			args: args{"0.1", "1.1"},
			want: -1,
		},
		{
			name: "5.4.3.2.1.00.00.0.05==5.000004.03.0002.1",
			args: args{"5.4.3.2.1.00.00.0.05", "5.000004.03.0002.1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
