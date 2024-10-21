package arrayList

import "testing"

var list *ArrayList[int]

func setUp() {
	list = NewArrayList[int](3)

}
func TestIsIfArrayListISEmpty(t *testing.T) {
	setUp()
	if !list.isEmpty() {
		t.Errorf("arrayList is suppose to be empty")
	}

}
func TestTheSizeOfTheArrayList(t *testing.T) {
	setUp()
	if list.Size() != 0 {
		t.Errorf("arrayList size is suppose to be 0")
	}

}
func TestTheAddMethodOfTheArrayList(t *testing.T) {
	setUp()
	intialSize := list.Size()
	list.add(1)
	list.add(2)

	if list.Size() != intialSize+2 {
		t.Errorf("arrayList size is suppose to be 3")
	}
}
