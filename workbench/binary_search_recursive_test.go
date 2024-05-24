package workbench

import "testing"

func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		list   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{[]int{1, 2, 3, 4}, 3}, 2},
		{"B", args{[]int{1, 2, 3, 4, 5}, 5}, 4},
		{"C", args{[]int{1, 2, 3}, 10}, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := BinarySearchRecursive(test.args.list, test.args.target); got != test.want {
				t.Errorf("BinarySearchRecursive() = %v, want %v", got, test.want)
			}
		})
	}
}
