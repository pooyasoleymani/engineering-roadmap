package sorting

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 3, 7, 3, 5, 7, 9, 1, 9}
	expct := []int{1, 1, 3, 3, 5, 7, 7, 9, 9}
	selectionSort(arr)
	if !slices.Equal(arr, expct) {
		t.Errorf("expect %v but got: %v", expct, arr)
	}
}
