package strings

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{""},
			want: "",
		},
		{
			name: "reverse",
			args: args{"reverse me"},
			want: "em esrever",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseAny(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "empty",
			args: args{nil},
			want: nil,
		},
		{
			name: "two strings",
			args: args{[]any{"one", "two"}},
			want: []any{"two", "one"},
		},
		{
			name: "slice of ints",
			args: args{[]any{5, 4, 3, 2, 1}},
			want: []any{1, 2, 3, 4, 5},
		},
		{
			name: "slice of runes",
			args: args{[]any{'a', 'b', 'c', 'd'}},
			want: []any{'d', 'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseAny(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
