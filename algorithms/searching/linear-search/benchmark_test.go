package serach

import "testing"

func BenchmarkLinearSerach(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for b.Loop() {
		LinearSearch(10, arr)
	}
}
