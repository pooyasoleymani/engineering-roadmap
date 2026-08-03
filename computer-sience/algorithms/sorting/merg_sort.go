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
	for i := range nr {
		leftArr[i] = arr[left+i]
	}
	for j := range nl {
		leftArr[j] = arr[mid+j]
	}

	i, j, k := 0, 0, left

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

func mergSort[T cmp.Ordered](arr []T, left, rigth int) {
	if left < rigth {
		mid := (left + rigth) / 2
		mergSort(arr, left, mid)
		mergSort(arr, mid+1, rigth)
		merge(arr, left, mid, rigth)
	}
}