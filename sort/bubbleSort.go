package sort

import (
	"dca0208/list"
)

// Sorts the given list through the bubbleSort implementation
func BubbleSort(arr list.List) {
	for j := 0; j < arr.Length()-1; j++ {
		ordered := true
		for i := 0; i < arr.Length()-1; i++ {
			if arr.Get(i) > arr.Get(i+1) {
				tmp := arr.Get(i)
				arr.Update(i, arr.Get(i+1))
				arr.Update(i+1, tmp)
				ordered = false
			}		
		}
		if ordered {return}
	}
}