package leetcode

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				chars: []byte("aabbccc"),
			},
			want: []byte("a2b2c3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compress(tt.args.chars)
			if got != len(tt.want) {
				t.Errorf("compress() = %d len, want %d len", got, len(tt.want))
			}
			if !reflect.DeepEqual(tt.args.chars[:got], tt.want) {
				t.Errorf("compress() = %v, want %v", tt.args.chars[:got], tt.want)
			}
		})
	}
}
