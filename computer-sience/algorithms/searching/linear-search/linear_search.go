package search

func LinearSearch(item int, arr []int) int {
	for _, i := range arr {
		if item == i {
			return item
		}
	}
	return -1
}
