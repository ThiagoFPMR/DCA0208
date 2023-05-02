package sort

import (
	"dca0208/list"
	"testing"
)

var implementations = map[string]func(arr list.List){
	"BubbleSort": BubbleSort,
}

var TestCases = map[string]struct {
	arrSize int
	arrData []int
}{
	"Empty": {0, []int{}},
	"Single": {1, []int{1}},
	"TwoOrdered": {2, []int{1, 2}},
	"TwoReversed": {2, []int{2, 1}},
	"SevenRandom": {7, []int{8, 4, 5, 3, 2, 9, 10}},
}

func TestSort(t *testing.T) {
	for testName, tt := range TestCases {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			arr := &list.ArrayList{}
			arr.Init()
			// Initialize array
			for i := 0; i < tt.arrSize; i++ {
				arr.AddToBack(tt.arrData[i])
			}

			for impName, implementation := range implementations {
				implementation := implementation
				t.Run(impName, func(t *testing.T) {
					t.Parallel()
					implementation(arr)
					for i := 0; i < arr.Length()-1; i++ {
						if arr.Get(i) > arr.Get(i+1) {
							t.Errorf("Does not behave as expected. List is not sorted. List Length: %d", arr.Length())
						}
					}
				})
			}
		})
			
	}
}