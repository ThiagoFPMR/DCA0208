package sort

import "dca0208/list"

// Sorts the given list through the countingSort implementation
func CountingSort(arr list.List) {
	max := arr.Get(0)
	min := arr.Get(0)

	for i := 1; i < arr.Length(); i++ {
		if arr.Get(i) > max {
			max = arr.Get(i)
		}
		if arr.Get(i) < min {
			min = arr.Get(i)
		}
	}

	countArr := make([]int, max-min+1)
	for i:=0; i < arr.Length(); i++ {
		countArr[arr.Get(i)-min]++
	}
	for i:=1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}

	outputArr := make([]int, arr.Length())
	for i:=0; i < arr.Length(); i++ {
		outputArr[countArr[arr.Get(i)-min]-1] = arr.Get(i)
		countArr[arr.Get(i)-min]--
	}

	// Copy the output array to the original array
	for i:=0; i < arr.Length(); i++ {
		arr.Update(i, outputArr[i])
	}
}