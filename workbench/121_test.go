package workbench

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{[]int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "B",
			args: args{[]int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "C",
			args: args{[]int{1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args.nums); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
