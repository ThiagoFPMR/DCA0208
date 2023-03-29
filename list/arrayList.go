package list

type ArrayList struct {
	values []int
	size   int
}

// Initialize the ArrayList with a default size of 10
func (arrayList *ArrayList) Init() {
	size := 10
	arrayList.values = make([]int, size)
	arrayList.size = 0
}

// Double the size of the ArrayList's underlying array
func (arrayList *ArrayList) double() {
	newValues := make([]int, cap(arrayList.values)*2)
	copy(newValues, arrayList.values)
	arrayList.values = newValues
}

// Return the length of the ArrayList
func (arrayList *ArrayList) Length() int {
	return arrayList.size
}

// Remove the last element of the ArrayList
func (arrayList *ArrayList) RemoveFromBack() {
	arrayList.values[arrayList.size-1] = 0
	arrayList.size--
}

// Add an element to the end of the ArrayList
func (arrayList *ArrayList) AddToBack(value int) {

	if cap(arrayList.values) == arrayList.size {
		arrayList.double()
	}

	arrayList.values[arrayList.size] = value

	arrayList.size++
}

// Add an element at the given index
func (arrayList *ArrayList) AddOnIndex(index int, value int) {
	// Make sure the ArrayList has enough space
	if cap(arrayList.values) == arrayList.size {
		arrayList.double()
	}

	for i := arrayList.size; i > index; i-- {
		arrayList.values[i] = arrayList.values[i-1]
	}
	arrayList.values[index] = value

	arrayList.size++
}

// Remove the element at the given index
func (arrayList *ArrayList) RemoveFromIndex(index int) {
	for i := index; i < arrayList.size; i++ {
		arrayList.values[i] = arrayList.values[i+1]
	}
	arrayList.size--
}

// Return the element at the given index
func (arrayList *ArrayList) Get(index int) int {
	return arrayList.values[index]
}

// Update the element at the given index
func (arrayList *ArrayList) Update(index int, value int) {
	arrayList.values[index] = value
}
