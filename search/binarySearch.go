package search

import (
	"dca0208/list"
	"fmt"
)

// BinarySearch returns the index of the target element in the given list.
func BinarySearch(arr *list.List, target, leftLimit, rightLimit int) int {

	mid := leftLimit + (rightLimit - leftLimit) / 2

	if leftLimit == rightLimit {
		return -1
	}

	fmt.Printf("leftLimit: %d, rightLimit: %d, mid: %d", leftLimit, rightLimit, mid)

	if (*arr).Get(mid) == target {
		return mid
	} else if (*arr).Get(mid) > target {
		return BinarySearch(arr, target, leftLimit, mid-1)
	} else {
		return BinarySearch(arr, target, mid+1, rightLimit)
	}
}