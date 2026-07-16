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

func (a *DynamicArray[T]) Len() int {
	return a.length
}

func (a *DynamicArray[T]) Capacity() int {
	return a.capacity
}

func (a *DynamicArray[T]) Push(value T) {
	if a.length == a.capacity {
		a.reAlloc()
	}
	a.elemets[a.length] = value
}

func (a *DynamicArray[T]) reAlloc() {
	a.capacity *= 2
	newArray := make([]T, a.capacity)
	copy(newArray, a.elemets)
	a.elemets = newArray

}

func (a *DynamicArray[T]) Insert(index int, value T) error {
	return nil
}

func (a *DynamicArray[T]) shiftRigth(index int) {
	length := a.length
	for i := length; i > index; i-- {
		if length == a.capacity {
			a.reAlloc()
		}
		
	}
}
