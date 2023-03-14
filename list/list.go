package list

type List interface {
	Init()
	Length() int
	RemoveFromBack()
	PushBack(value int)
	Get(index int) int
	Remove(index int, value int)
	Insert(index int, value int)
	Update(index int, value int)
}