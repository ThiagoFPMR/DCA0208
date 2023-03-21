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
