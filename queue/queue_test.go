package queue

import (
	"testing"
)

// Useful when you only want to run a subset of tests.
var tests = map[string]func(t *testing.T){
	"Length":  testLength,
	"Peek":    testPeek,
	"Enqueue": testEnqueue,
	"Dequeue": testDequeue,
	"IsEmpty": testIsEmpty,
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
	queue := new(Queue)
	queue.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		queue.Enqueue(i)
	}

	if queue.Length() != limit {
		t.Errorf("Length of instance is %d, expected %d", queue.Length(), limit)
	}
}

var peekTests = map[string]struct {
	frontValue, queueSize int
}{
	"NonEmpty": {15, 15},
	"Empty":    {15, 0},
}

// Tests the Enqueue method.
func testPeek(t *testing.T) {
	for name, tt := range peekTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			queue := new(Queue)
			queue.Init()

			for i := 0; i < tt.queueSize; i++ {
				queue.Enqueue(tt.frontValue)
			}

			value, error := queue.Peek()

			// Tests relating to the value returned by Pop
			if value != tt.frontValue && tt.queueSize > 0 {
				t.Errorf("Peek returned %d, expected %d", value, tt.frontValue)
			}
			if tt.queueSize == 0 && value != 0 {
				t.Errorf("Peek returned %d, expected %d", value, 0)
			}

			// Tests relating to the size after Peek
			if tt.queueSize != queue.Length() {
				t.Errorf("Peek altered the queue's size. It should not have. Size is %d, expected %d", queue.Length(), tt.queueSize)
			}

			// Tests relating to the error returned by Peek
			if error != nil && tt.queueSize > 0 {
				t.Errorf("Peek returned error, expected no error")
			}
			if error == nil && tt.queueSize == 0 {
				t.Errorf("Peek returned no error, expected error")
			}
		})
	}
}

// Tests the Enqueue method.
func testEnqueue(t *testing.T) {
	queue := new(Queue)
	queue.Init()

	limit := 15
	for i := 0; i < limit; i++ {
		queue.Enqueue(i)
	}

	if queue.Length() != limit {
		t.Errorf("Length of instance is %d, expected %d", queue.Length(), limit)
	}

	for i := 0; i < limit-1; i++ {
		queue.Dequeue()
	}

	value, _ := queue.Peek()
	if value != limit-1 {
		t.Errorf("Value at the front of the queue does not match the inserted value. The value is %d, expected %d", value, limit-1)
	}
}

var dequeueTests = map[string]struct {
	frontValue, queueSize int
}{
	"Empty":    {3, 0},
	"NonEmpty": {27, 15},
}

// Tests the Dequeue method.
func testDequeue(t *testing.T) {
	for name, tt := range dequeueTests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			queue := new(Queue)
			queue.Init()

			for i := 0; i < tt.queueSize; i++ {
				queue.Enqueue(tt.frontValue)
			}

			value, error := queue.Dequeue()

			// Tests relating to the value returned by Dequeue
			if value != tt.frontValue && tt.queueSize > 0 {
				t.Errorf("Dequeue of instance is %d, expected %d", value, tt.frontValue)
			}
			if tt.queueSize == 0 && value != 0 {
				t.Errorf("Dequeue of instance is %d, expected %d", value, 0)
			}

			// Tests relating to the size after Dequeue
			if tt.queueSize > 0 && queue.Length() != tt.queueSize-1 {
				t.Errorf("Failed to dequeue element properly. Length of instance is %d, expected %d", queue.Length(), tt.queueSize-1)
			}
			if tt.queueSize == 0 && queue.Length() != tt.queueSize {
				t.Errorf("Attempted to dequeue an empty queue.")
			}

			// Tests relating to the error returned by Dequeue
			if error != nil && tt.queueSize > 0 {
				t.Errorf("Dequeue returned error, expected no error")
			}
			if error == nil && tt.queueSize == 0 {
				t.Errorf("Dequeue didn't return error, expected error")
			}
		})
	}
}

var isEmptyTests = map[string]struct {
	queueSize int
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
