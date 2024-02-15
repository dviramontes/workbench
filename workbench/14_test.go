package workbench

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{[]string{"flow", "flower", "flight"}},
			want: "fl",
		},
		{
			name: "B",
			args: args{[]string{"rat", "racecar", "rare", "rap"}},
			want: "ra",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
