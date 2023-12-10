package utils

type Stack[T any] []*T

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	} else {
		i := len(*s) - 1
		val := (*s)[i]
		*s = (*s)[:i]

		return val
	}
}

func (s *Stack[T]) Push(value *T) {
	*s = append(*s, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
