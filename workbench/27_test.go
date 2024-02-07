package workbench

import (
	"reflect"
	"testing"
)

func TestRemoveElementFromArray(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name        string
		args        args
		wantNexNums int
	}{
		{
			name:        "A",
			args:        args{[]int{3, 2, 2, 3}, 3},
			wantNexNums: 2,
		},
		{
			name: "B",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			wantNexNums: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNexNums := RemoveElementFromArray(tt.args.nums, tt.args.val); !reflect.DeepEqual(gotNexNums, tt.wantNexNums) {
				t.Errorf("RemoveElementFromArray() => %v, want => %v", gotNexNums, tt.wantNexNums)
			}
		})
	}
}
