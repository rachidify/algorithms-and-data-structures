package search

import "math"

// BinarySearch ...
func BinarySearch(list []int, key int) int {
	var mid, elm int
	low := 0
	high := len(list) - 1

	for low <= high {
		mid = int(math.Floor(float64((low + high) / 2)))
		elm = list[mid]

		if elm < key {
			low = mid - 1
		} else if elm > key {
			high = mid + 1
		} else {
			return mid
		}

	}

	return -1
}
