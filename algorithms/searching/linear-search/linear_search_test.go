package serach

import "testing"

func TestLinearSearchExist(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := LinearSearch(10, arr)
	if got != 10 {
		t.Fatalf("expect: 10 but got: %v", got)
	}

}

func TestLinearSearchNotExist(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	got := LinearSearch(-5, arr)
	if got != -1 {
		t.Fatalf("expect: -1 but got: %v", got)
	}
}
