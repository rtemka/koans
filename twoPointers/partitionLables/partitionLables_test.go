package leetcode

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "1",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partitionLabels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
