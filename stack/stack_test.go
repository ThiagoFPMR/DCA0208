package stack

import (
	"testing"
)

// Useful when you only want to run a subset of tests.
var tests = map[string]func(t *testing.T){
	"Length": testLength,
	"IsEmpty": testIsEmpty,
	"Push": testPush,
	"Peek": testPeek,
	"Pop": testPop,
}

// Runs all tests of the Stack interface.
func TestStack(t *testing.T) {
	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			test(t)
		})
	}
}

// Tests the Length method.
func testLength(t *testing.T) {
	stack := new(Stack)
	stack.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		stack.Push(i)
	}

	if stack.Length() != limit {
		t.Errorf("Length of instance is %d, expected %d", stack.Length(), limit)
	}
}

var isEmptyTests = map[string]struct {
	stackSize int
}{
	"Empty": {0},
	"NonEmpty": {15},
}


// Tests the IsEmpty method.
func testIsEmpty(t *testing.T) {
	for name, tt := range isEmptyTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stack := new(Stack)
			stack.Init()

			for i := 0; i < tt.stackSize; i++ {
				stack.Push(i)
			}

			if stack.IsEmpty() != (tt.stackSize == 0) {
				t.Errorf("IsEmpty returned %t, expected %t", stack.IsEmpty(), tt.stackSize == 0)
			}
		})
	}
}

// Tests the Push method.
func testPush(t *testing.T) {
	stack := new(Stack)
	stack.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		stack.Push(i)
	}

	if stack.Length() != limit {
		t.Errorf("Failed to push element. Length of instance is %d, expected %d", stack.Length(), limit)
	}

	if stack.Peek() != limit-1 {
		t.Errorf("Failed to push element. Element at the top is %d, expected %d", stack.Peek(), limit-1)
	}
}

var popTests = map[string]struct {
	topElement, stackSize int
}{
	"NonEmpty": {15, 15},
	"Empty": {15, 0},
}

// Tests the Pop method.
func testPop(t *testing.T) {
	for name, tt := range popTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stack := new(Stack)
			stack.Init()

			for i := 0; i < tt.stackSize; i++ {
				stack.Push(tt.topElement)
			}
			
			value := stack.Pop()

			if tt.stackSize == 0 {
				if value != 0 {
					t.Errorf("Pop returned %d, expected %d", value, 0)
				}
				if stack.Length() != 0 {
					t.Errorf("Failed to pop element. Length of instance is %d, expected %d", stack.Length(), 0)
				}
				return
			} else {
				if value != tt.topElement {
					t.Errorf("Pop returned %d, expected %d", value, tt.topElement)
				}
			}
			if stack.Length() != tt.stackSize - 1 {
				t.Errorf("Failed to pop element. Length of instance is %d, expected %d", stack.Length(), tt.stackSize - 1)
			}
		})
	}
}

// Tests the Peek method.
func testPeek(t *testing.T) {
	stack := new(Stack)
	stack.Init()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Peek() != 3 {
		t.Errorf("Peek returned %d, expected %d", stack.Peek(), 3)
	}
}



