package workbench

import "testing"

func TestValidParenthesis(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A",
			args: args{"()[]{}"},
			want: true,
		},
		{
			name: "B",
			args: args{"[)"},
			want: false,
		},
		{
			name: "C",
			args: args{"ABC"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParenthesis(tt.args.s); got != tt.want {
				t.Errorf("ValidParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
