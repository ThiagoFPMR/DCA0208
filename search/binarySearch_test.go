package search

import (
	"dca0208/list"
	"testing"
)

var implementations = map[string]func(arr list.List, target int) int{
	"RecursiveBinarySearch": RecursiveBinarySearch,
}

var binarySearchTests = map[string]struct {
	arr                     list.List
	arrSize, target, output int
}{
	"OutOfBounds":   {&list.LinkedList{}, 10, 10, -1},
	"ValueAtStart":  {&list.LinkedList{}, 10, 0, 0},
	"ValueAtEnd":    {&list.LinkedList{}, 10, 9, 9},
	"ValueAtMiddle": {&list.LinkedList{}, 10, 5, 5},
}

func TestBinarySearch(t *testing.T) {
	for testName, tt := range binarySearchTests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			tt.arr.Init()

			for i := 0; i < tt.arrSize; i++ {
				tt.arr.AddToBack(i)
			}

			for impName, implementation := range implementations {
				impName := impName
				implementation := implementation
				t.Run(impName, func(t *testing.T) {
					t.Parallel()
					result := implementation(tt.arr, tt.target)
					if result != tt.output {
						t.Errorf("Does not behave as expected. Returned %d, expected %d. Target: %d. List Length: %d", result, tt.output, tt.target, tt.arr.Length())
					}
				})
			}
		})
	}
}
