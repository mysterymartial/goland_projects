package arrayList

type ArrayList[T any] struct {
	size     int
	capacity int
	elements []T
}

func NewArrayList[T any](capacity int) *ArrayList[T] {
	return &ArrayList[T]{
		capacity: capacity,
		elements: make([]T, capacity),
		size:     0,
	}

}
func (l ArrayList[T]) isEmpty() bool {
	return l.size == 0

}

func (l ArrayList[T]) Size() int {
	return l.size
}

func (l ArrayList[T]) add(element T) {
	if l.size == l.capacity {
		l.ensureCapacity()

	}
	l.elements[l.size] = element
	l.size++
}

func (l ArrayList[T]) ensureCapacity() {
	if l.size == l.capacity {
		l.capacity *= 2
		new_elements := make([]T, l.capacity)
		copy(new_elements, l.elements)
		l.elements = new_elements

	}

}
