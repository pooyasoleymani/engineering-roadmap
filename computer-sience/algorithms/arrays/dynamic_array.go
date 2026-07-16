package arrays

type DynamicArray[T any] struct {
	length   int
	capacity int
	elemets  []T
}

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &DynamicArray[T]{
		length:   0,
		capacity: capacity,
		elemets:  make([]T, capacity),
	}
}

func (a *DynamicArray[T]) Capacity() int {
	return a.capacity
}

func (a *DynamicArray[T]) Push(value T) {
	if a.length == a.capacity {
		a.reAlloc()
	}
	a.elemets[a.length] = value
	a.length++
}

func (a *DynamicArray[T]) reAlloc() {
	a.capacity *= 2
	newArray := make([]T, a.capacity)
	for i, element := range a.elemets {
		newArray[i] = element
	}
	a.elemets = newArray

}

func (a *DynamicArray[T]) Insert(index int, value T) {
	a.shiftRigth(index)
	a.elemets[index] = value
	a.length++
}

func (a *DynamicArray[T]) shiftRigth(index int) {
	length := a.length
	for i := length; i > index; i-- {
		if length == a.capacity {
			a.reAlloc()
		}
		a.elemets[i] = a.elemets[i-1]
	}
}

func (a *DynamicArray[T]) shiftLeft(index int) {
	for i := index; i < a.length; i++ {
		a.elemets[i] = a.elemets[i+1]
	}
	a.length--
}

func (a *DynamicArray[T]) Pop() T {
	a.length--
	return a.elemets[a.length-1]
}

func (a *DynamicArray[T]) Set(index int, value T) {
	a.elemets[index] = value
}

func (a *DynamicArray[T]) Size() int {
	return a.length
}

func (a *DynamicArray[T]) Remove(index int) {
	a.shiftLeft(index)
	a.length--
}
