package workbench

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		list   []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"A", args{[]int{4,1,3,2}}, []int{1,2,3,4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := QuickSort(test.args.list); !reflect.DeepEqual(got, test.want) {
				t.Errorf("QuickSort() = %v, want %v", got, test.want)
			}
		})
	}
}
