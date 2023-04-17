package list

type doubleNode struct {
	value int
	next  *doubleNode
	prev  *doubleNode
}

type DoublyLinkedList struct {
	head *doubleNode
	tail *doubleNode
	size int
}

// Initialize the DoublyLinkedList.
func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = &doubleNode{value: 0, next: nil, prev: nil}
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
		doublyLinkedList.head = &doubleNode{value: value, next: nil, prev: nil}
		doublyLinkedList.tail = doublyLinkedList.head
	} else {
		doublyLinkedList.tail.next = &doubleNode{value: value, next: nil, prev: doublyLinkedList.tail}
		doublyLinkedList.tail = doublyLinkedList.tail.next
	}

	doublyLinkedList.size++
}

// Remove the last element of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) RemoveFromBack() {
	if doublyLinkedList.size == 1 {
		doublyLinkedList.head = &doubleNode{value: 0, next: nil, prev: nil}
		doublyLinkedList.tail = doublyLinkedList.head
	} else {
		doublyLinkedList.tail.prev.next = nil
		doublyLinkedList.tail = doublyLinkedList.tail.prev
	}	

	doublyLinkedList.size--
}

// Add an element at the given index
func (doublyLinkedList *DoublyLinkedList) AddOnIndex(index int, value int) {
	var current *doubleNode = doublyLinkedList.head

	if index == 0 {
		doublyLinkedList.head = &doubleNode{value: value, next: current, prev: nil}
		current.prev = doublyLinkedList.head
		doublyLinkedList.size++
		return
	}

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	next := current.next

	newdoubleNode := &doubleNode{value: value, next: current.next, prev: current}
	current.next = newdoubleNode

	if next != nil {
		next.prev = newdoubleNode
	}
	if next == nil {
		doublyLinkedList.tail = newdoubleNode
	}

	doublyLinkedList.size++
}

// Remove an element at the given index
func (doublyLinkedList *DoublyLinkedList) RemoveFromIndex(index int) {
	var current *doubleNode = doublyLinkedList.head

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
	var current *doubleNode = doublyLinkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

// Update the element at the given index
func (doublyLinkedList *DoublyLinkedList) Update(index int, value int) {
	var current *doubleNode = doublyLinkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = value
}
