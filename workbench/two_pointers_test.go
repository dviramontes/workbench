package workbench

import (
	"reflect"
	"testing"
)

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

func TestCombineArrayUsingTwoPointers(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name  string
		args  args
		wantC []int
	}{
		{
			name: "A",
			args: args{
				a: []int{1, 4, 7, 20},
				b: []int{3, 5, 6, 8},
			},
			wantC: []int{1, 3, 4, 5, 6, 7, 8, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CombineArrayUsingTwoPointers(tt.args.a, tt.args.b); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("CombineArrayUsingTwoPointers() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
