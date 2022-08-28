package leetcode

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{[]string{}},
			want: nil,
		},
		{
			name: "two_of_four",
			args: args{[]string{"Hello", "AlaSka", "Dad", "Peace"}},
			want: []string{"AlaSka", "Dad"},
		},
		{
			name: "one_wrong",
			args: args{[]string{"omk"}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
