package queue

import "errors"

type Queue[T any] struct {
	size     int
	capacity int
	elements []T
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		capacity: capacity,
		elements: make([]T, 0, capacity),
		size:     0,
	}

}
func (q *Queue[T]) isEmpty() bool {
	return q.size == 0

}

func (q *Queue[T]) Size() int {
	return q.size

}

func (q *Queue[T]) enqueue(element T) {
	if q.size == q.capacity {
		q.ensureCapacity()
	}
	q.elements = append(q.elements, element)
	q.size++

}

func (q *Queue[T]) ensureCapacity() {
	q.capacity *= 2
	new_elements := make([]T, q.capacity)
	copy(new_elements, q.elements)
	q.elements = new_elements

}

func (q *Queue[T]) dequeue() (T, error) {
	var zero T
	if q.isEmpty() {
		return zero, errors.New("queue is empty")
	}
	element := q.elements[0]
	copy(q.elements[:0], q.elements[1:q.size])
	q.elements[q.size-1] = zero
	q.size--
	return element, nil

}

func (q *Queue[T]) peek() (T, error) {
	var zero T
	if q.isEmpty() {
		return zero, errors.New("queue is empty")

	}
	return q.elements[0], nil

}
func (q *Queue[T]) getCapacity() int {
	return q.capacity

}
