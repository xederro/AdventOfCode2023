package utils

type PriorityQueue[T any] struct {
	queue   []*T
	compare func(a, b *T) int
}

func (q *PriorityQueue[T]) SetComparator(comp func(a, b *T) int) {
	q.compare = comp
}

func (q *PriorityQueue[T]) Dequeue() *T {
	val := 0
	if q.IsEmpty() {
		return nil
	}
	for k := range q.queue {
		if q.compare(q.queue[val], q.queue[k]) >= 0 {
			val = k
		} else {

		}
	}
	r := q.queue[val]
	q.queue = append(q.queue[:val], q.queue[val+1:]...)
	return r
}

func (q *PriorityQueue[T]) Enqueue(value *T) {
	q.queue = append(q.queue, value)
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}
