package arrays

import (
	"testing"
)

func TestBuiltinCopyToInsert(t *testing.T) {
	arr := NewDynamicArray[float32](0)
	arr.Insert(0, 1.)
	arr.Insert(0, 2.)
	arr.Insert(1, 3.)

	capacity := arr.Capacity()

	if capacity != 4 {
		t.Errorf("expect: %d but got: %d", 4, capacity)
	}

	if size := arr.Size(); size != 3 {
		t.Errorf("expect: %d but got: %d", 3, size)
	}

	if element := arr.Get(1); element != 3.0 {
		t.Errorf("expect: %f but got: %f", 3.0, element)
	}
}
