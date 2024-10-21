package stack

import "errors"

type Stack[T any] struct {
	size     int
	capacity int
	elements []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		size:     0,
		capacity: capacity,
		elements: make([]T, capacity),
	}

}
func (s *Stack[T]) isEmpty() bool {
	return s.size == 0
}
func (s *Stack[T]) Size() int {
	return s.size

}

func (s *Stack[T]) push(element T) {
	if s.size == s.capacity {
		s.ensureCapacity()
	}
	s.elements[s.size] = element
	s.size++

}

func (s *Stack[T]) ensureCapacity() {
	s.capacity *= 2
	new_elements := make([]T, s.capacity)
	copy(new_elements, s.elements)
	s.elements = new_elements
}

func (s *Stack[T]) pop() (T, error) {
	var zeroval T
	if s.isEmpty() {
		return zeroval, errors.New("stack is empty")
	}
	s.size--
	element := s.elements[s.size]
	s.elements[s.size] = zeroval
	return element, nil
}

func (s *Stack[T]) peek() (T, error) {
	var zeroval T
	if s.isEmpty() {
		return zeroval, errors.New("stack is empty")
	}
	return s.elements[s.size-1], nil

}

func (s *Stack[T]) GetCapacity() int {
	return s.capacity

}
