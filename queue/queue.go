package queue

type Queue interface {
	Init()
	Enqueue(value int) 
	Dequeue() (int, error)
	Peek() (int, error)
	Length() int
	IsEmpty() bool
}