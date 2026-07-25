package genericbinarysearch

import (
	"cmp"
	"testing"
)

type tests[T cmp.Ordered] = []struct {
	arr    []T
	target T
	value  T
}

func ptr[T any](v T) *T {
	return &v
}

func TestLowerBuond(t *testing.T) {
	testCases := tests[int]{
		{[]int{1, 2, 3, 3}, 3, 3},
		{[]int{1}, 1, 1},
	}

	for _, test := range testCases {

		got := LowerBound(test.target, test.arr)

		if got != test.value {
			t.Errorf("expect: %v but got: %v", test.value, got)
		}
	}
}

func TestUpperBuond(t *testing.T) {
	testCases := tests[int]{
		{[]int{1, 2, 3, 3, 5, 7, 9}, 3, 5},
		{[]int{1}, 1, 1},
	}

	for _, test := range testCases {

		got := UpperBound(test.target, test.arr)

		if got != test.value {
			t.Errorf("expect: %v but got: %v", test.value, got)
		}
	}
}
