package workbench

import (
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		items []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		wantOp bool
	}{
		{
			name: "A",
			fields: fields{
				items: []int{1, 2, 3},
			},
			want:   1,
			wantOp: true,
		},
		{
			name: "B",
			fields: fields{
				items: []int{},
			},
			want:   0,
			wantOp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				items: tt.fields.items,
			}
			gotValue, gotError := q.Dequeue()
			if gotValue != tt.want {
				t.Errorf("Dequeue() got = %v, want-val %v", gotValue, tt.want)
			}
			if gotError != tt.wantOp {
				t.Errorf("Dequeue() got1 = %v, want-op %v", gotError, tt.wantOp)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		items []int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
	}{
		{
			name: "A",
			fields: fields{
				items: []int{1, 2, 3},
			},
			args: args{
				num: 4,
			},
			wantSize: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				items: tt.fields.items,
			}
			q.Enqueue(tt.args.num)
			if q.Size() != tt.wantSize {
				t.Errorf("Enqueue(), want-size %v", tt.wantSize)
			}
		})
	}
}
