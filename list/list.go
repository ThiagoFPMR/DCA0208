package list

type List interface {
	Init()
	Length() int
	RemoveFromBack()
	AddToBack(value int)
	AddOnIndex(index int, value int)
	RemoveFromIndex(index int)
	Get(index int) int
	Update(index int, value int)
}
