package workbench

import "testing"

func TestTwoSumUsingTwoPointers(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A",
			args: args{[]int{1, 2, 4, 6, 8, 9, 14, 15}, 13},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumUsingTwoPointers(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("TwoSumUsingTwoPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
