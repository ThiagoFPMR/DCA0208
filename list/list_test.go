package list

import (
	"testing"
)

/*
	The purpose of structuring the code like this is to make it easy to add new
	implementations of the List interface as well as new tests. All while keeping
	the code DRY.
*/

// All implementations of the List interface.
var instances = map[string]List{
	"ArrayList":        &ArrayList{},
	"LinkedList":       &LinkedList{},
	"DoublyLinkedList": &DoublyLinkedList{},
}

// All tests to run on all implementations of the List interface.
var tests = map[string]func(t *testing.T, instance List){
	"TestLength":          testLength,
	"TestAddToBack":       testAddToBack,
	"TestRemoveFromBack":  testRemoveFromBack,
	"TestAddOnIndex":      testAddOnIndex,
	"TestRemoveFromIndex": testRemoveFromIndex,
	"TestGet":             testGet,
	"TestUpdate":          testUpdate,
}

// TestLists runs all tests on all implementations of the List interface.
func TestLists(t *testing.T) {
	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			for instName, instance := range instances {
				instance := instance
				t.Run(instName, func(t *testing.T) {
					t.Parallel()
					test(t, instance)
				})
			}
		})
	}
}

// Functions for testing the List interface.

// Tests the Length method.
func testLength(t *testing.T, instance List) {
	instance.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		instance.AddToBack(i)
	}

	if instance.Length() != limit {
		t.Errorf("Length of instance is %d, expected %d", instance.Length(), limit)
	}
}

// Tests the AddToBack method.
func testAddToBack(t *testing.T, instance List) {
	instance.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		instance.AddToBack(i)
	}

	if instance.Length() != limit {
		t.Errorf("Failed to add element. Length of instance is %d, expected %d", instance.Length(), limit)
	}

	if instance.Get(limit-1) != limit-1 {
		t.Errorf("Failed to add element. Element at index %d is %d, expected %d", limit-1, instance.Get(limit-1), limit-1)
	}
}

// Tests the RemoveFromBack method.
func testRemoveFromBack(t *testing.T, instance List) {
	instance.Init()

	instance.AddToBack(1)
	instance.RemoveFromBack()

	if instance.Length() > 0 {
		t.Errorf("Failed to remove element. Length of instance is %d, expected %d", instance.Length(), 0)
	}
}

var AddOnIndexTests = map[string]struct {
	startingSize, index, value int
}{
	"AddToBeggining": {15, 0, 1},
	"AddOToMiddle":   {13, 7, 2},
	"AddToEnd":       {11, 11, 3},
}

// Tests the AddOnIndex method.
func testAddOnIndex(t *testing.T, instance List) {
	for name, tt := range AddOnIndexTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			instance.Init()

			limit := tt.startingSize
			for i := 0; i < limit; i++ {
				instance.AddToBack(i)
			}

			instance.AddOnIndex(tt.index, tt.value)

			if instance.Get(tt.index) != tt.value {
				t.Errorf("Failed to add element on index %d. Element value is %d, expected %d", tt.index, instance.Get(tt.index), tt.value)
			}

			if instance.Length() != limit+1 {
				t.Errorf("Failed to add element. Length of instance is %d, expected %d", instance.Length(), limit+1)
			}
		})
	}
}

var RemoveFromIndexTests = map[string]struct {
	startingSize, index int
}{
	"RemoveFromBeggining": {15, 0},
	"RemoveFromMiddle":    {13, 7},
	"RemoveFromEnd":       {11, 11},
}

// Tests the RemoveFromIndex method.
func testRemoveFromIndex(t *testing.T, instance List) {
	for name, tt := range RemoveFromIndexTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			instance.Init()

			limit := tt.startingSize
			for i := 0; i < limit; i++ {
				instance.AddToBack(i)
			}

			nextElement := 0
			if tt.index < limit-1 {
				nextElement = instance.Get(tt.index + 1)
			}

			instance.RemoveFromIndex(tt.index)

			if instance.Length() > 1 && tt.index < limit-1 {
				if instance.Get(tt.index) != nextElement {
					t.Errorf("The element that replaced the removed element is %d, expected %d", instance.Get(tt.index), nextElement)
				}
			}

			if instance.Length() != limit-1 {
				t.Errorf("Failed to add element. Length of instance is %d, expected %d", instance.Length(), limit-1)
			}
		})
	}
}

var GetTests = map[string]struct {
	startingSize, index int
}{
	"ValidIndex": {15, 1},
	//"InvalidIndex": {13, 15},
}

// Tests the Get method.
func testGet(t *testing.T, instance List) {
	for name, tt := range GetTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			instance.Init()

			limit := tt.startingSize
			for i := 0; i < limit; i++ {
				instance.AddToBack(i)
			}

			if tt.index < limit {
				if instance.Get(tt.index) != tt.index {
					t.Errorf("Failed to get element on index %d. Element at index %d is %d, expected %d", tt.index, tt.index, instance.Get(tt.index), tt.index)
				}
			} else {
				return
			}
		})
	}
}

// Tests the Update method.
func testUpdate(t *testing.T, instance List) {
	instance.Init()

	instance.AddToBack(1)
	instance.AddToBack(2)

	instance.Update(1, 3)

	if instance.Get(1) != 3 {
		t.Errorf("Failed to update element on index 1. Element at index 1 is %d, expected %d", instance.Get(1), 3)
	}
}
