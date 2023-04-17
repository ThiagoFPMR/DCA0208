package tree

type Tree interface {
	Init()
	Height() int
	IsEmpty() bool
	Add(value int)
	Search(value int) bool
	Remove(value int) bool
}
