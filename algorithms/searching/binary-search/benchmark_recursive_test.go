package search

import "testing"

func BenchmarBinarySerachRecursive(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for b.Loop() {
		BinarySearchRecursive(10, 0, len(arr), arr)
	}
}
