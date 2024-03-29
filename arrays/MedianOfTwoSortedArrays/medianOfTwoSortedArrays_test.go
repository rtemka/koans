package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "[1,3]+[2]=2.0",
			args: args{nums1: []int{1, 3}, nums2: []int{2}},
			want: 2.0,
		},
		{
			name: "[1,2]+[3,4]=2.5",
			args: args{nums1: []int{1, 2}, nums2: []int{3, 4}},
			want: 2.5,
		},
		{
			name: "[1,2,5,6]+[3,4]=3.5",
			args: args{nums1: []int{1, 2, 5, 6}, nums2: []int{3, 4}},
			want: 3.5,
		},
		{
			name: "[1,3]+[2,7]=2.5",
			args: args{nums1: []int{1, 3}, nums2: []int{2, 7}},
			want: 2.5,
		},
		{
			name: "[3]+[-2,-1]=-1.0",
			args: args{nums1: []int{3}, nums2: []int{-2, -1}},
			want: -1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
