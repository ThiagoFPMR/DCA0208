package stack

import (
	"testing"
)

// Useful when you only want to run a subset of tests.
var tests = map[string]func(t *testing.T){
	"Length":  testLength,
	"IsEmpty": testIsEmpty,
	"Push":    testPush,
	"Peek":    testPeek,
	"Pop":     testPop,
}

// Runs all tests of the Stack interface.
func TestStack(t *testing.T) {
	for testName, test := range tests {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
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
	"Empty":    {0},
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

	value, _ := stack.Peek()
	if value != limit-1 {
		t.Errorf("Failed to push element. Element at the top is %d, expected %d", value, limit-1)
	}
}

var popTests = map[string]struct {
	topValue, stackSize int
}{
	"NonEmpty": {15, 15},
	"Empty":    {15, 0},
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
				stack.Push(tt.topValue)
			}

			value, error := stack.Pop()

			// Tests relating to the value returned by Pop
			if value != tt.topValue && tt.stackSize > 0 {
				t.Errorf("Pop returned %d, expected %d", value, tt.topValue)
			}
			if tt.stackSize == 0 && value != 0 {
				t.Errorf("Pop returned %d, expected %d", value, 0)
			}

			// Tests relating to the size after Pop
			if tt.stackSize > 0 && stack.Length() != tt.stackSize-1 {
				t.Errorf("Failed to pop element. Length of instance is %d, expected %d", stack.Length(), tt.stackSize-1)
			}
			if tt.stackSize == 0 && stack.Length() != tt.stackSize {
				t.Errorf("Attempted to pop an empty stack.")
			}

			// Tests relating to the error returned by Pop
			if error != nil && tt.stackSize > 0 {
				t.Errorf("Pop returned error, expected no error")
			}
			if error == nil && tt.stackSize == 0 {
				t.Errorf("Pop returned no error, expected error")
			}
		})
	}
}

var peekTests = map[string]struct {
	topValue, stackSize int
}{
	"NonEmpty": {15, 15},
	"Empty":    {15, 0},
}

// Tests the Peek method.
func testPeek(t *testing.T) {
	for name, tt := range peekTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stack := new(Stack)
			stack.Init()

			for i := 0; i < tt.stackSize; i++ {
				stack.Push(tt.topValue)
			}

			value, error := stack.Peek()

			// Tests relating to the value returned by Pop
			if value != tt.topValue && tt.stackSize > 0 {
				t.Errorf("Peek returned %d, expected %d", value, tt.topValue)
			}
			if tt.stackSize == 0 && value != 0 {
				t.Errorf("Peek returned %d, expected %d", value, 0)
			}

			// Tests relating to the size after Pop
			if tt.stackSize != stack.Length() {
				t.Errorf("Peek altered the stack's size. It should not have. Size is %d, expected %d", stack.Length(), tt.stackSize)
			}

			// Tests relating to the error returned by Peek
			if error != nil && tt.stackSize > 0 {
				t.Errorf("Peek returned error, expected no error")
			}
			if error == nil && tt.stackSize == 0 {
				t.Errorf("Peek returned no error, expected error")
			}
		})
	}
}
