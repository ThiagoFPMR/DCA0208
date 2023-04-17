package queue

import (
	"errors"
)

type node struct {
	value int
	next  *node
}

type Queue struct {
	front *node
	size  int
}

// Initialize the Queue.
func (queue *Queue) Init() {
	queue.front = nil
	queue.size = 0
}

// Return the length of the Queue
func (queue *Queue) Length() int {
	return queue.size
}

// Check if the Queue is empty
func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

// Add an element to the back of the Queue
func (queue *Queue) Enqueue(value int) {
	if queue.IsEmpty() {
		queue.front = &node{value: value, next: nil}
	} else {
		var tail *node = queue.front
		for i := 0; i < queue.size-1; i++ {
			tail = tail.next
		}
		tail.next = &node{value: value, next: nil}
	}
	queue.size++
}

// Remove the front element of the Queue
func (queue *Queue) Dequeue() (int, error) {
	if !queue.IsEmpty() {
		value := queue.front.value
		if queue.size == 1 {
			queue.front = nil
		} else {
			queue.front = queue.front.next
		}
		queue.size--
		return value, nil
	}
	return 0, errors.New("Queue is empty")
}

// Peek the front element of the Queue
func (queue *Queue) Peek() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	} else {
		return queue.front.value, nil
	}
}
