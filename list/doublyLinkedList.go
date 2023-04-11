package list

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}

type DoublyLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

// Initialize the DoublyLinkedList.
func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = &DoubleNode{value: 0, next: nil, prev: nil}
	doublyLinkedList.tail = doublyLinkedList.head
	doublyLinkedList.size = 0
}

// Return the length of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) Length() int {
	return doublyLinkedList.size
}

// Add an element to the end of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) AddToBack(value int) {

	if doublyLinkedList.size == 0 {
		doublyLinkedList.head = &DoubleNode{value: value, next: nil, prev: nil}
		doublyLinkedList.tail = doublyLinkedList.head
	} else {
		doublyLinkedList.tail.next = &DoubleNode{value: value, next: nil, prev: doublyLinkedList.tail}
		doublyLinkedList.tail = doublyLinkedList.tail.next
	}

	doublyLinkedList.size++
}

// Remove the last element of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) RemoveFromBack() {
	if doublyLinkedList.size == 1 {
		doublyLinkedList.head = &DoubleNode{value: 0, next: nil, prev: nil}
		doublyLinkedList.tail = doublyLinkedList.head
	} else {
		doublyLinkedList.tail.prev.next = nil
		doublyLinkedList.tail = doublyLinkedList.tail.prev
	}	

	doublyLinkedList.size--
}

// Add an element at the given index
func (doublyLinkedList *DoublyLinkedList) AddOnIndex(index int, value int) {
	var current *DoubleNode = doublyLinkedList.head

	if index == 0 {
		doublyLinkedList.head = &DoubleNode{value: value, next: current, prev: nil}
		current.prev = doublyLinkedList.head
		doublyLinkedList.size++
		return
	}

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	next := current.next

	newDoubleNode := &DoubleNode{value: value, next: current.next, prev: current}
	current.next = newDoubleNode

	if next != nil {
		next.prev = newDoubleNode
	}
	if next == nil {
		doublyLinkedList.tail = newDoubleNode
	}

	doublyLinkedList.size++
}

// Remove an element at the given index
func (doublyLinkedList *DoublyLinkedList) RemoveFromIndex(index int) {
	var current *DoubleNode = doublyLinkedList.head

	if index == 0 {
		doublyLinkedList.head = current.next
		doublyLinkedList.head.prev = nil
		doublyLinkedList.size--
		return
	}

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	if index < doublyLinkedList.size-1 {
		current.next = current.next.next
		current.next.prev = current
	} else {
		current.next = nil
		doublyLinkedList.tail = current
	}

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
