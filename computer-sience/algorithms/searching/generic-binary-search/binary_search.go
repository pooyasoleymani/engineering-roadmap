package genericbinarysearch

import (
	"cmp"
	"sort"
)

// Search returns the target if found, otherwise returns the zero value of T.
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
		} else {
			low = mid + 1
		}
	}
	sort.Search()
	return res
}

// LowerBound returns the index of the first element >= target.
// Returns len(arr) if all elements are strictly less than target.
func LowerBound[T cmp.Ordered](target T, arr []T) int {
	res := len(arr) // Default fallback if no elements are >= target
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] >= target {
			res = mid      // Candidate found, try to find a smaller index
			high = mid - 1 // Shift left
		} else {
			low = mid + 1 // Shift right
		}
	}
	return res
}

// UpperBound returns the index of the first element > target.
// Returns len(arr) if all elements are less than or equal to target.
func UpperBound[T cmp.Ordered](target T, arr []T) int {
	res := len(arr) // Default fallback if no elements are > target
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] > target {
			res = mid      // Candidate found, try to find a smaller index
			high = mid - 1 // Shift left
		} else {
			low = mid + 1 // Shift right
		}
	}
	return res
}

// EqualBound returns the index of the first exact occurrence of target.
// Returns -1 if the target is not present in the array.
func EqualBound[T cmp.Ordered](target T, arr []T) int {
	res := -1
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			res = mid
			high = mid - 1 // Shift left to find the absolute *first* occurrence
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
