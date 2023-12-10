package utils

type Queue[T any] []*T

func (q *Queue[T]) Dequeue() *T {
	val := (*q)[0]
	if !q.IsEmpty() {
		*q = (*q)[1:]
	}

	return val
}

func (q *Queue[T]) Enqueue(value *T) {
	*q = append(*q, value)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
