package sort

import "dca0208/list"

func partition(arr list.List, leftLimit, rightLimit int) int {
	return 1
}

func quickSort(arr list.List, leftLimit, rightLimit int) {
	if leftLimit < rightLimit {
		pivot := partition(arr, leftLimit, rightLimit)
		quickSort(arr, leftLimit, pivot-1) // left side
		quickSort(arr, pivot+1, rightLimit) // right side
	}
}

func QuickSort(arr list.List) {
	quickSort(arr, 0, arr.Length()-1)
}