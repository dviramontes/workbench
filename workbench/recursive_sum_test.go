package workbench

import "testing"

func TestRecursiveSum(t *testing.T) {
	type args struct {
		list []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{[]int{1, 2, 3, 4, 5}}, 15},
		{"B", args{[]int{8, 9, 10}}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecursiveSum(tt.args.list); got != tt.want {
				t.Errorf("RecursiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
