package workbench

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: true,
		},
		{
			name: "single character",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "two characters",
			args: args{s: "aa"},
			want: true,
		},
		{
			name: "two different characters",
			args: args{s: "ab"},
			want: false,
		},
		{
			name: "three characters",
			args: args{s: "aba"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
