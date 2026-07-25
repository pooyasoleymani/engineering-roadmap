package genericbinarysearch

import "cmp"

func Search[T cmp.Ordered](target T, arr []T) T {
	var res T
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return target
		} else if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		}
	}
	return res
}

func LowerBound[T cmp.Ordered](target T, arr []T) T {
	var res T
	low := 0
	high := len(arr) - 1

	for low <= high && res < target {
		mid := low + (high-low)/2
		if target == arr[mid] {
			res = target
		} else if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			res = arr[mid]
			low = mid + 1
		}
	}
	return res
}

func UpperBound[T cmp.Ordered](target T, arr []T) T {
	var res T
	low := 0
	high := len(arr) - 1

	for low <= high && res <= target {
		mid := low + (high-low)/2
		if target == arr[mid] {
			res = target
		} else if target < arr[mid] {
			res = arr[mid]
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		}
	}
	return res
}

func EquealBound[T cmp.Ordered](target T, arr []T) T {
	var res T
	low := 0
	high := len(arr) - 1

	for low <= high || res != target {
		mid := low + (high-low)/2
		if target == arr[mid] {
			res = target
		} else if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		}
	}
	return res
}
