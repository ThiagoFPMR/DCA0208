package sort

import "dca0208/list"

func merge(arr, leftArr, rightArr list.List) {
	arrIndex := 0
	leftIndex := 0
	rightIndex := 0

	for leftIndex < leftArr.Length() && rightIndex < rightArr.Length() {
		if leftArr.Get(leftIndex) <= rightArr.Get(rightIndex) {
			arr.Update(arrIndex, leftArr.Get(leftIndex))
			leftIndex++
		} else {
			arr.Update(arrIndex, rightArr.Get(rightIndex))
			rightIndex++
		}
		arrIndex++
	}

	for leftIndex < leftArr.Length() {
		arr.Update(arrIndex, leftArr.Get(leftIndex))
		leftIndex++
		arrIndex++
	}

	for rightIndex < rightArr.Length() {
		arr.Update(arrIndex, rightArr.Get(rightIndex))
		rightIndex++
		arrIndex++
	}
}

func MergeSort(arr list.List) {
	if arr.Length() <= 1 {
		return
	}
	mid := arr.Length() / 2

	leftArr := &list.ArrayList{}
	leftArr.Init()
	for i := 0; i < mid; i++ {
		leftArr.AddToBack(arr.Get(i))
	}

	rightArr := &list.ArrayList{}
	rightArr.Init()
	for i := mid; i < arr.Length(); i++ {
		rightArr.AddToBack(arr.Get(i))
	}

	MergeSort(leftArr)
	MergeSort(rightArr)

	merge(arr, leftArr, rightArr)
}
