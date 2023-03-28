package list

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

// Initialize the LinkedList.
func (linkedList *LinkedList) Init() {
	linkedList.head = &Node{value: 0, next: nil}
	linkedList.size = 0
}

// Return the length of the LinkedList
func (linkedList *LinkedList) Length() int {
	return linkedList.size
}

// Add an element to the end of the LinkedList
func (linkedList *LinkedList) AddToBack(value int) {
	var tail *Node = linkedList.head

	for i := 0; i < linkedList.size-1; i++ {
		tail = tail.next
	}

	tail.next = &Node{value: value, next: nil}

	linkedList.size++
}

// Remove the last element of the LinkedList
func (linkedList *LinkedList) RemoveFromBack() {
	var tail *Node = linkedList.head

	for i := 0; i < linkedList.size-1; i++ {
		tail = tail.next
	}
	tail.next = nil

	linkedList.size--
}

// Add an element at the given index
func (linkedList *LinkedList) AddOnIndex(index int, value int) {
	var current *Node = linkedList.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = &Node{value: value, next: current.next}

	linkedList.size++
}

// Remove the element at the given index
func (linkedList *LinkedList) RemoveFromIndex(index int) {
	var current *Node = linkedList.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}
	current.next = current.next.next

	linkedList.size--
}

// Return the element at the given index
func (linkedList *LinkedList) Get(index int) int {
	var current *Node = linkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

// Update the element at the given index
func (linkedList *LinkedList) Update(index int, value int) {
	var current *Node = linkedList.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = value
}
