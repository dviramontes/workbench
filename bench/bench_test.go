package bench

import (
	"reflect"
	"testing"
)

func Test_maxMeetings(t *testing.T) {
	type args struct {
		start []int
		end   []int
		want  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"A", args{
			start: []int{1, 3, 0, 5, 8, 5},
			end:   []int{2, 4, 6, 7, 9, 9},
		}, []int{1, 2, 4, 5}},
		//{"B", args{
		//	start: []int{75250, 50074, 43659, 8931, 11273, 27545, 50879, 77924},
		//	end:   []int{112960, 114515, 81825, 93424, 54316, 35533, 73383, 160252},
		//}, []int{6, 7, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxMeetings(tt.args.start, tt.args.end)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("MaxMeetings() = %v, want %v", MaxMeetings(tt.args.start, tt.args.end), tt.want)
			}
		})
	}
}

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"A", args{
				intervals: []Interval{{5, 10}, {0, 30}, {15, 20}},
			}, false,
		},
		{
			"B", args{
				intervals: []Interval{{5, 8}, {9, 15}},
			}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("CanAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeAlternate(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}

	testing.Verbose()

	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{
				word1: "abc",
				word2: "xyz",
			},
			want: "axbycz",
		},
		{
			name: "B",
			args: args{
				word1: "abcd",
				word2: "pq",
			},
			want: "apbqcd",
		},
		{
			name: "C",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "D",
			args: args{
				word1: "cdf",
				word2: "a",
			},
			want: "cadf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAlternate(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MergeAlternate() = %v, want %v", got, tt.want)
			}
		})
	}
}
