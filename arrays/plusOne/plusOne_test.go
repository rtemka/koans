package plusone

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{digits: []int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.args.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("plusOne() = %v, want %v", got, tt.want)
			}

		})
	}

}
