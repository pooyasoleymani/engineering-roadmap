package search

func BinarySearchRecursive(item int, low int, high int, arr []int) int {
	mid := low + (high-low)/2
	if item == arr[mid] {
		return item
	} else if item < arr[mid] {
		high = mid - 1
		return BinarySearchRecursive(item, low, high, arr)
	} else if item > arr[mid] {
		low = mid + 1
		return BinarySearchRecursive(item, low, high, arr)
	}
	return -1
}
