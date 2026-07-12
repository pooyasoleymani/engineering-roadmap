package search

import "testing"

func TestBinarySearchRecursivExist(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := BinarySearchRecursive(9, 0, len(arr), arr)
	if got != 9 {
		t.Errorf("expect: %v but got: %v", 9, got)
	}
}

func TestBinarySearchRecursivNotExist(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := BinarySearchRecursive(-5, 0, len(arr), arr)
	if got != -1 {
		t.Errorf("expect: %v but got: %v", -1, got)
	}
}
