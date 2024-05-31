package workbench

import "testing"

func TestLinkedList(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Insert 1",
			args: args{
				data: 1,
			},
		},
		{
			name: "Search",
			args: args{
				data: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := &LinkedList{}
			ll.Insert(tt.args.data)
			if ll.length != 1 {
				t.Errorf("Insert() = %v, want %v", ll.length, 1)
			}

			for v := range 10 {
				ll.Insert(v)
			}

			if ll.length != 11 {
				t.Errorf("Insert() = %v, want %v", ll.length, 11)
			}

			ll.Insert(4)

			if result := ll.Search(4); result {
				t.Errorf("Search() = %v, want %v", 4, true)
			}
		})
	}
}
