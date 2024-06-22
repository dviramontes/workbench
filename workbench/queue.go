package workbench

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{items: make([]int, 0)}
}

func (q *Queue) Enqueue(num int) {
	q.items = append(q.items, num)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
