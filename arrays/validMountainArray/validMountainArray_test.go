package validmountainarray

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if validMountainArray(tt.args.arr) != tt.want {
				t.Fatal("validMountainArray() = wrong answer")
			}
		})
	}
}
