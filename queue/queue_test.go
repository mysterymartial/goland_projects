package queue

import "testing"

var queue *Queue[int]

func setup() {
	queue = NewQueue[int](3)

}
func TestIfQueueIsEmpty(t *testing.T) {
	setup()
	if !queue.isEmpty() {
		t.Error("Queue is not empty")
	}

}
func TestTheSizeOfQueue(t *testing.T) {
	setup()
	if queue.Size() != 0 {
		t.Error("Queue size is suppose to 0 when empty")
	}

}
func TestEnqueueFunctionality(t *testing.T) {
	setup()
	intialSize := queue.Size()
	queue.enqueue(1)
	queue.enqueue(2)
	if queue.Size() != intialSize+2 {
		t.Errorf("Queue size is suppose to 2  ")
	}

}
func TestDequeueFunctionality(t *testing.T) {
	setup()
	queue.enqueue(1)
	queue.enqueue(2)
	val, err := queue.dequeue()
	if err != nil {
		t.Errorf("Error dequeuing element")
	}
	if val != 1 {
		t.Errorf("element dequed is suppose to 1  ")
	}
	if queue.Size() != 1 {
		t.Errorf("Queue size is suppose to 1  ")
	}

}
func TestPopingWhileQueueIsEmpty(t *testing.T) {
	setup()
	_, err := queue.dequeue()
	if err == nil {
		t.Errorf("you cant deque from an empty queue")

	}

}
func TestPeekFunctionality(t *testing.T) {
	setup()
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	val, err := queue.peek()
	if err != nil {
		t.Errorf("Error peeking element")
	}
	if val != 1 {
		t.Errorf("element peeked is suppose to 1  ")
	}

}
func TestEnsureCapacity(t *testing.T) {
	smallQueue := NewQueue[int](1)
	smallQueue.enqueue(1)
	smallQueue.enqueue(2)
	smallQueue.enqueue(3)
	val, err := smallQueue.peek()
	if err != nil {
		t.Errorf("Error peeking element")
	}
	if val != 1 {
		t.Errorf("element peeked is suppose to 1  ")
	}
	if smallQueue.getCapacity() <= 1 {
		t.Errorf("Queue capacity is supposed to increase  ")
	}

}
