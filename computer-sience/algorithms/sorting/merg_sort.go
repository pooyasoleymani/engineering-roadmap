package sorting

import (
	"cmp"
)

func merge[T cmp.Ordered](arr []T, left, mid, rigth int) {
	nl := mid - left + 1
	nr := rigth - mid

	// create temp arrays
	leftArr := make([]T, nl)
	rigthArr := make([]T, nr)

	// copy data to temp arrays
	for i := range nl {
		leftArr[i] = arr[left+i]
	}

	for j := range nr {
		leftArr[j] = arr[mid+j+1]
	}

	i := 0
	j := 0
	k := left

	// merge the temp array back
	for i < nl && j < nr {
		if leftArr[i] <= rigthArr[j] {
			arr[k] = leftArr[i]
			i++

		} else {
			arr[k] = rigthArr[j]
			j++
		}
		k++
	}

	// copy remaining elements of leftArr
	for i < nl {
		arr[k] = leftArr[i]
		i++
		k++
	}
	// copy rmainig elements of rigthArr
	for j < nr {
		arr[k] = rigthArr[j]
		j++
		k++
	}
}

func mergSort[T cmp.Ordered](arr []T, left, right int) {
	if left < right {
		mid := left + (right-left)/2 // Prevents potential integer overflow
		mergSort(arr, left, mid)
		mergSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}
