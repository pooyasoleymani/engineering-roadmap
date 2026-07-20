package arrays

import (
	"container/list"
	"testing"
)

func BenchmarkInsertLinkedList(b *testing.B) {
	ll := list.New()
	for b.Loop() {
		ll.PushBack(1)
	}
}

func BenchmarkInsertArray(b *testing.B) {
	arr := make([]int, 0)
	for b.Loop() {
		arr = append(arr, 1)
	}
}

func BenchmarkIterateArray(b *testing.B) {
	var result int
	arr := make([]int, 12121, 1_000_000)

	b.ResetTimer()

	for b.Loop() {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		result = sum
	}
	_ = result
}

func BenchmarkIterateLinkedList(b *testing.B) {
	var result int
	ll := list.New()
	arr := make([]int, 12129, 1_000_000)
	for _, a := range arr {
		ll.PushBack(a)
	}
	b.ResetTimer()

	for b.Loop() {
		sum := 0
		for e := ll.Front(); e != nil; e = e.Next() {
			sum += e.Value.(int)
		}
		result = sum
	}
	_ = result

}
