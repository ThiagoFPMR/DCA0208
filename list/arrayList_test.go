package list

import (
	"testing"
)

// TestLength tests the Length method of the ArrayList
func TestLength(t *testing.T) {
	var arrayList ArrayList
	arrayList.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		arrayList.AddToBack(i)
	}

	if arrayList.Length() != limit {
		t.Errorf("Length of ArrayList is %d, expected %d", arrayList.Length(), limit)
	}
}

func TestRemoveFromBack(t *testing.T) {
	var arrayList ArrayList
	arrayList.Init()

	arrayList.AddToBack(i)
	arrayList.RemoveFromBack()

	if arrayList.Length() > 0 {
		t.Errorf("Failed to remove element. Length of ArrayList is %d, expected %d", arrayList.Length(), limit-1)
	}
}

func TestAddOnIndex(t *testing.T) {
	var arrayList ArrayList
	arrayList.Init()

	arrayList.AddToBack(1)
	arrayList.AddToBack(2)
	arrayList.AddToBack(3)
	arrayList.AddOnIndex(1, 4)

	if arrayList.Get(1) != 4 {
		t.Errorf("Failed to add element on index 1. Element at index 1 is %d, expected %d", arrayList.Get(1), 4)
	}
}