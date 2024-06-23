package workbench

import (
	"testing"
)

func TestSetSize(t *testing.T) {
	type args struct {
		list []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{[]int{1, 1, 2, 2, 3, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet()
			for _, v := range tt.args.list {
				s.Add(v)
			}
			got := s.Size()
			if got != tt.want {
				t.Errorf("Set()/size = %v, want-size = %v", got, tt.want)
			}
		})
	}
}
