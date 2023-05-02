package sort

import "dca0208/list"

// Sorts the given list through the insertionSort implementation
func InsertionSort(arr list.List) {
	for i := 1; i < arr.Length(); i++ {
		for j := i; j > 0 && arr.Get(j) < arr.Get(j-1); j-- {
			tmp := arr.Get(j)
			arr.Update(j, arr.Get(j-1))
			arr.Update(j-1, tmp)
		}
	}
}