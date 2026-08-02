package sorting

import (
	"cmp"
)

func mergSort[T cmp.Ordered](arr []T) {
	left := arr[:len(arr)/2]
	rigth := arr[len(arr)/2:]

	if len(left) == 0 {
		return
	}

	
	mergSort(rigth)
	mergSort(left)
}
