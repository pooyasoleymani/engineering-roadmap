package genericbinarysearch

import "cmp"

func Search[T cmp.Ordered](item T, arr []T) T {
	var res T
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if item == arr[mid] {
			return item
		} else if item < arr[mid] {
			high = mid - 1
		} else if item > arr[mid] {
			low = mid + 1
		}
	}
	return res
}
