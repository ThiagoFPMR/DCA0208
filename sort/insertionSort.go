package sort

import "dca0208/list"

// Sorts the given list through the insertionSort implementation
func InsertionSort(arr list.List) {
	for i := 1; i < arr.Length(); i++ {
		value := arr.Get(i)
		var j int
		for j = i; j > 0 && value < arr.Get(j-1); j-- {
			arr.Update(j, arr.Get(j-1))
		}
		arr.Update(j, value)
	}
}
