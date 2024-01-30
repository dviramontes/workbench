package workbench

import "testing"

func TestNumberOfDiceRollsToTarget(t *testing.T) {
	type args struct {
		d      int
		k      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{1, 6, 3}, 1},
		{"B", args{2, 6, 7}, 6},
		{"C", args{30, 30, 500}, 222616187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfDiceRollsToTarget(tt.args.d, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("NumberOfDiceRollsToTarget() got => %v, want => %v", got, tt.want)
			}
		})
	}
}
