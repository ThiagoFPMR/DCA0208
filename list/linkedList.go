package list

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
	size int
}

// Initialize the LinkedList.
func (linkedList *LinkedList) Init() {
	linkedList.head = nil
	linkedList.size = 0
}

// Return the length of the LinkedList
func (linkedList *LinkedList) Length() int {
	return linkedList.size
}

// Add an element to the end of the LinkedList
func (linkedList *LinkedList) AddToBack(value int) {
	if linkedList.size == 0 {
		linkedList.head = &node{value: value, next: nil}
	} else {
		var tail *node = linkedList.head
		for i := 0; i < linkedList.size-1; i++ {
			tail = tail.next
		}
		tail.next = &node{value: value, next: nil}
	}
	linkedList.size++
}

// Remove the last element of the LinkedList
func (linkedList *LinkedList) RemoveFromBack() {
	if linkedList.size == 1 {
		linkedList.head = nil
		linkedList.size--
		return
	}

	var tail *node = linkedList.head
	for i := 0; i < linkedList.size-1; i++ {
		tail = tail.next
	}
	tail.next = nil

	linkedList.size--
}

// Add an element at the given index
func (linkedList *LinkedList) AddOnIndex(index int, value int) {
	var current *node = linkedList.head

	if index == 0 {
		linkedList.head = &node{value: value, next: current}
		linkedList.size++
		return
	}

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = &node{value: value, next: current.next}

	linkedList.size++
}

// Remove the element at the given index
func (linkedList *LinkedList) RemoveFromIndex(index int) {
	var current *node = linkedList.head

	if index == 0 {
		linkedList.head = current.next
		linkedList.size--
		return
	}

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	if index < linkedList.size-1 {
		current.next = current.next.next
	} else {
		current.next = nil
	}

	linkedList.size--
}

// Return the element at the given index
func (linkedList *LinkedList) Get(index int) int {
	var current *node = linkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

// Update the element at the given index
func (linkedList *LinkedList) Update(index int, value int) {
	var current *node = linkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = value
}
