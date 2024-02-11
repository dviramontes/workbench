package workbench

import (
	"reflect"
	"testing"
)

func TestRotateArrayNTimes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "A",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "B",
			args: args{[]int{-1, -100, 3, 99}, 2},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "C",
			args: args{[]int{1, 2, 3, 4}, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "D",
			args: args{[]int{1, 2, 3, 4}, 2},
			want: []int{3, 4, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateArrayNTimes(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateArrayNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
