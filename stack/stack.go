package stack

import (
	"errors"
)

type StackNode struct {
	value int
	prev  *StackNode
}

type Stack struct {
	top  *StackNode
	size int
}

// Initialize the Stack.
func (stack *Stack) Init() {
	stack.top = nil
	stack.size = 0
}

// Return the length of the Stack
func (stack *Stack) Length() int {
	return stack.size
}

// Check if the Stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

// Add an element to the top of the Stack
func (stack *Stack) Push(value int) {
	stack.top = &StackNode{value: value, prev: stack.top}
	stack.size++
}

// Peek the top element of the Stack
func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	} else {
		return stack.top.value, nil
	}
}

// Remove the top element of the Stack
func (stack *Stack) Pop() (int, error) {
	if !stack.IsEmpty() {
		value := stack.top.value
		stack.top = stack.top.prev
		stack.size--
		return value, nil
	}
	return 0, errors.New("Stack is empty")
}


