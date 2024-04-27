package workbench

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A",
			args: args{
				list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 5,
			},
			want: 4,
		},
		{
			name: "B",
			args: args{
				list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
