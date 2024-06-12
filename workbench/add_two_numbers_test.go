package workbench

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *Node
		l2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "A",
			args: args{
				l1: &Node{Data: 2, Next: &Node{Data: 4, Next: &Node{Data: 3, Next: nil}}},
				l2: &Node{Data: 5, Next: &Node{Data: 6, Next: &Node{Data: 4, Next: nil}}},
			},
			want: &Node{Data: 7, Next: &Node{Data: 0, Next: &Node{Data: 8, Next: nil}}},
		},
		{
			name: "B",
			args: args{
				l1: &Node{Data: 0, Next: nil},
				l2: &Node{Data: 0, Next: nil},
			},
			want: &Node{Data: 0, Next: nil},
		},
		{
			name: "C",
			args: args{
				l1: &Node{Data: 9, Next: &Node{Data: 9, Next: &Node{Data: 9, Next: &Node{Data: 9, Next: nil}}}},
				l2: &Node{Data: 9, Next: &Node{Data: 9, Next: &Node{Data: 9, Next: &Node{Data: 9, Next: nil}}}},
			},
			want: &Node{
				Data: 8,
				Next: &Node{
					Data: 9,
					Next: &Node{
						Data: 9,
						Next: &Node{
							Data: 9,
							Next: &Node{
								Data: 1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
