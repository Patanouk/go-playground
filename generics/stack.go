package generics

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	numElements := len(s.elements) - 1
	result := s.elements[numElements]
	s.elements = s.elements[:numElements]
	return result, true
}
