package queue

import (
	"testing"
)

// Useful when you only want to run a subset of tests.
var tests = map[string]func(t *testing.T){
	"Length": testLength,
}

// Runs all tests of the Stack interface.
func TestQueue(t *testing.T) {
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
	queue := *new(Queue)
	queue.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		queue.Enqueue(i)
	}

	if queue.Length() != limit {
		t.Errorf("Length of instance is %d, expected %d", queue.Length(), limit)
	}
}

// Tests the Enqueue method.
func testPeek(t *testing.T) {
	queue := *new(Queue)
	queue.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		queue.Enqueue(i)
	}

	if queue.Peek() != 0 {
		t.Errorf("Peek of instance is %d, expected %d", queue.Peek(), 0)
	}
}

var dequeueTests = map[string]struct {
	frontValue, queueSize int
}{
	"Empty": {3, 0},
	"NonEmpty": {27, 15},
}

// Tests the Dequeue method.
func testDequeue(t *testing.T) {
	for name, tt := range dequeueTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			queue := *new(Queue)
			queue.Init()

			for i := 0; i < tt.queueSize; i++ {
				queue.Enqueue(i)
			}

			value, error:= queue.Dequeue()

			// Tests relating to the value returned by Dequeue
			if value != tt.frontValue && tt.queueSize > 0 {
				t.Errorf("Dequeue of instance is %d, expected %d", value, tt.frontValue)
			}
			if tt.queueSize == 0 && value != 0 {
				t.Errorf("Dequeue of instance is %d, expected %d", value, 0)
			}

			// Tests relating too the size after Dequeue
			if tt.queueSize != 0 && queue.Length() != tt.queueSize - 1 {
				t.Errorf("Failed to dequeue element properly. Length of instance is %d, expected %d", queue.Length(), tt.queueSize - 1)
			}
			if tt.queueSize != 0 && queue.Length() != tt.queueSize {
				t.Errorf("Attempted to dequeue an empty queue.")
			}

			// Tests relating to the error returned by Dequeue
			if error == nil && tt.queueSize > 0 {
				t.Errorf("Dequeue returned error, expected no error")
			}
			if error != nil && tt.queueSize == 0 {
				t.Errorf("Dequeue didn't return error, expected error")
			}
		})
	}
}

var isEmptyTests = map[string]struct {
	queueSize int
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
			stack := *new(Queue)
			stack.Init()

			for i := 0; i < tt.queueSize; i++ {
				stack.Enqueue(i)
			}

			if stack.IsEmpty() != (tt.queueSize == 0) {
				t.Errorf("IsEmpty returned %t, expected %t", stack.IsEmpty(), tt.queueSize == 0)
			}
		})
	}
}


