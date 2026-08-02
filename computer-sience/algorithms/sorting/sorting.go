package sorting

import "cmp"

func selectionSort[T cmp.Ordered](arr []T) {
	n := len(arr)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func bubbleSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		p := i - 1
		for p > 0 && arr[p+1] < arr[p] {
			arr[p], arr[p+1] = arr[p+1], arr[p]
		}
	}
}
