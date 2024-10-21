package stack

import "testing"

var stack *Stack[int]

func setup() {
	stack = NewStack[int](3)
}
func TestNewStackIsEmpty(t *testing.T) {
	setup()
	if !stack.isEmpty() {
		t.Errorf("stack is not empty")
	}

}
func TestSizeOfStack(t *testing.T) {
	setup()
	if stack.Size() != 0 {
		t.Errorf("stack size at intial is more than 0")

	}
}
func TestPushRefelctOnSize(t *testing.T) {
	setup()
	intialSize := stack.Size()
	stack.push(1)
	stack.push(2)
	if stack.Size() != intialSize+2 {
		t.Errorf("stack size is not equal to 2")
	}
}

func TestPopReflectOnSize(t *testing.T) {
	setup()
	stack.push(1)
	stack.push(2)
	stack.pop()
	val, err := stack.pop()
	if err != nil {
		t.Errorf("stack pop fail")
	}
	if val != 1 {
		t.Errorf("stack size is not equal to 1")
	}

}
func TestPopingAnEmptyStack(t *testing.T) {
	setup()
	_, err := stack.pop()
	if err == nil {
		t.Errorf("Expected an error when poping from the stack")
	}

}
func TestForThePeekFunctionality(t *testing.T) {
	setup()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	val, err := stack.peek()
	if err != nil {
		t.Errorf("stack peek fail")
	}
	if val != 3 {
		t.Errorf("stack element that should peek the last element")
	}
}
func TestThatTheCapacityCanIncrease(t *testing.T) {

	smallStack := NewStack[int](1)
	smallStack.push(1)
	smallStack.push(2)
	smallStack.push(3)
	val, err := smallStack.peek()
	if err != nil {
		t.Errorf("stack peek fail")
	}
	if val != 3 {
		t.Errorf("stack element that should peek the last element")
	}
	if smallStack.GetCapacity() <= 1 {
		t.Errorf("expected stack capacity to increase")
	}
}
