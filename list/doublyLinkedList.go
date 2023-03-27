package list

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	size int
}

// Initialize the DoublyLinkedList.
func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = &DoubleNode{value: 0, next: nil, prev: nil}
	doublyLinkedList.size = 0
}

// Return the length of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) Length() int {
	return doublyLinkedList.size
}

// Add an element to the end of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) AddToBack(value int) {
	var tail *DoubleNode = doublyLinkedList.head

	for i := 0; i < doublyLinkedList.size-1; i++ {
		tail = tail.next
	}

	tail.next = &DoubleNode{value: value, next: nil, prev: tail}

	doublyLinkedList.size++
}

// Remove the last element of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) RemoveFromBack() {
	var tail *DoubleNode = doublyLinkedList.head

	for i := 0; i < doublyLinkedList.size-1; i++ {
		tail = tail.next
	}
	tail.next = nil

	doublyLinkedList.size--
}

// Add an element at the given index
func (doublyLinkedList *DoublyLinkedList) AddOnIndex(index int, value int) {
	var current *DoubleNode = doublyLinkedList.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	next := current.next

	newDoubleNode := &DoubleNode{value: value, next: current.next, prev: current}
	current.next = newDoubleNode

	if next != nil {
		next.prev = newDoubleNode
	}

	doublyLinkedList.size++
}

// Remove an element at the given index
func (doublyLinkedList *DoublyLinkedList) RemoveOnIndex(index int) {
	var current *DoubleNode = doublyLinkedList.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = current.next.next
	current.next.prev = current

	doublyLinkedList.size--
}

// Get the value of the element at the given index
func (doublyLinkedList *DoublyLinkedList) Get(index int) int {
	var current *DoubleNode = doublyLinkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

// Update the element at the given index
func (doublyLinkedList *DoublyLinkedList) Update(index int, value int) {
	var current *DoubleNode = doublyLinkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = value
}