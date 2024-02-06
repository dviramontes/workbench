package workbench

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "A",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			wantResult: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "B",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			wantResult: []int{1},
		},
		{
			name: "C",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			wantResult: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MergeSortedArray(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MergeSortedArray() => %v, want => %v", gotResult, tt.wantResult)
			}
		})
	}
}
