package list

import (
	"testing"
)

var instances = map[string]List {
	"instance" : &ArrayList{},
	"LinkedList" : &LinkedList{},
	"DoublyLinkedList" : &DoublyLinkedList{},
}


func TestLength(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()	

			limit := 15
			for i := 0; i < limit; i++ {
				instance.AddToBack(i)
			}

			if instance.Length() != limit {
				t.Errorf("Length of instance is %d, expected %d", instance.Length(), limit)
			}
		})
	}
}

func TestAddToBack(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
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
		})
	}
}


func TestRemoveFromBack(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()

			instance.AddToBack(1)
			instance.RemoveFromBack()

			if instance.Length() > 0 {
				t.Errorf("Failed to remove element. Length of instance is %d, expected %d", instance.Length(), 0)
			}
		})
	}
}


func TestAddOnIndex(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()

			instance.AddToBack(1)
			instance.AddToBack(2)
			instance.AddToBack(3)

			previousValue := instance.Get(1)

			instance.AddOnIndex(1, 4)

			if instance.Get(1) != 4 {
				t.Errorf("Failed to add element on index 1. Element at index 1 is %d, expected %d", instance.Get(1), 4)
			}

			if instance.Get(2) != previousValue {
				t.Errorf("Failed to move elements to the right during insetion. Element at next index is %d, expected %d", instance.Get(2), previousValue)
			}
		})
	}
}


func TestRemoveOnIndex(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()

			instance.AddToBack(1)
			instance.AddToBack(2)
			instance.AddToBack(3)

			replacementValue := instance.Get(2)

			instance.RemoveFromIndex(1)

			if instance.Get(1) != replacementValue {
				t.Errorf("Failed to remove element on index 1. Element at index 1 is %d, expected %d", instance.Get(1), 3)
			}
		})
	}
}

func TestGet(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()

			instance.AddToBack(1)
			instance.AddToBack(2)

			if instance.Get(1) != 2 {
				t.Errorf("Failed to get element on index 1. Element at index 1 is %d, expected %d", instance.Get(1), 2)
			}
		})
	}
}


func TestUpdate(t *testing.T) {
	for name, instance := range instances {
		instance := instance
		t.Run(name, func(t *testing.T) {
			instance.Init()

			instance.AddToBack(1)
			instance.AddToBack(2)

			instance.Update(1, 3)

			if instance.Get(1) != 3 {
				t.Errorf("Failed to update element on index 1. Element at index 1 is %d, expected %d", instance.Get(1), 3)
			}
		})
	}
}
