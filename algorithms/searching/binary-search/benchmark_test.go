package search

import "testing"

func BenchmarBinarySerach(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for b.Loop() {
		BinarySearch(10, arr)
	}
}
