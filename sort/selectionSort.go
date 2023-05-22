package sort

import "dca0208/list"

// Sorts the given list through the selectionSort implementation
func SelectionSort(arr list.List) {
	for i := 0; i < arr.Length()-1; i++ {
		min := i
		for j := i+1; j < arr.Length(); j++ {
			if arr.Get(j) < arr.Get(min) {
				min = j
			}
		}
		tmp := arr.Get(i)
		arr.Update(i, arr.Get(min))
		arr.Update(min, tmp)
	}
}