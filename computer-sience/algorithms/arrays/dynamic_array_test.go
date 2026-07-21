package arrays

import (
	"algorithm/arrays"
	"testing"
)

func TestBuiltinCopyToInsert(t *testing.T) {
	arr := arrays.NewDynamicArray[float64](3)
	arr.Insert(1.2323)
	arr.Insert(1.1212)
	arr.Insert(2.2323)

	capacity := arr.Capacity()

	if capacity != 3 {
		t.Errorf("expect: %d but got: %d", 3, capacity)
	}

	if size := arr.Size(); size != 3 {
		t.Errorf("expect: %d but got: %d", 3, size)
	}

}
