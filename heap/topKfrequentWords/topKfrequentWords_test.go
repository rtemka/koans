package leetcode

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{
			name:  "1",
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     2,
			want:  []string{"i", "love"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.words, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
