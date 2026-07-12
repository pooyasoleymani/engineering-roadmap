package search

func BinarySearch(item int, arr []int) int {
	length := len(arr)
	low := 0
	high := length - 1

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
	return -1
}
