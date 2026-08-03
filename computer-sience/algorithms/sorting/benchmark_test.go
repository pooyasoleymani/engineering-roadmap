package sorting

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(size * 10)
	}
	return slice
}

func BenchmarkMergeSort(b *testing.B) {
	size := 10000 // Test with 10,000 elements
	baseData := generateRandomSlice(size)
	b.ResetTimer()

	for b.Loop() {
		data := make([]int, len(baseData))
		copy(data, baseData)
		mergSort(data, 0, len(data)-1)
	}
}

func BenchmarkBuiltInSort(b *testing.B) {
	size := 10000 // Test with 10,000 elements
	baseData := generateRandomSlice(size)
	b.ResetTimer()

	for b.Loop() {
		data := make([]int, len(baseData))
		copy(data, baseData)
		slices.SortStableFunc(data, func(a, b int) int { return cmp.Compare(a, b) })
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	size := 10000 // Test with 10,000 elements
	baseData := generateRandomSlice(size)
	b.ResetTimer()

	for b.Loop() {
		data := make([]int, len(baseData))
		copy(data, baseData)
		insertionSort(data)
	}
}
