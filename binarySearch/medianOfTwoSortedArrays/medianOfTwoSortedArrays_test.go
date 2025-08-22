package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
