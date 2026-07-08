package serach

import "testing"

func TestLinearSearch(t *testing.T) {
	arrays := []struct{
		[]int{1, 2, 3, 4, 5, 6, 7, 8 ,9, 10},
		[]int{-5, -4, -3, 1, 5, 8, 10}
	}
	for arr := range arrays {
		got := LinearSearch(10, arr)
		if got != 10 {
			t.Fatalf("expect: 10 but got: %v", got)
		}
	}

}