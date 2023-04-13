package search

import (
	"dca0208/list"
)

// Recursive binary search implementation
func recursiveBinarySearch(arr list.List, target, leftLimit, rightLimit int) int {

	mid := leftLimit + (rightLimit-leftLimit)/2

	if arr.Get(mid) == target {
		return mid
	}

	if leftLimit == rightLimit {
		return -1
	}

	if arr.Get(mid) > target {
		return recursiveBinarySearch(arr, target, leftLimit, mid-1)
	} else {
		return recursiveBinarySearch(arr, target, mid+1, rightLimit)
	}
}

// Returns the index of the target element in the array. Returns -1 if the element is not found.
// RecursiveBinarySearch is a wrapper for recursiveBinarySearch, which properly implements the method.
func RecursiveBinarySearch(arr list.List, target int) int {
	return recursiveBinarySearch(arr, target, 0, arr.Length()-1)
}
