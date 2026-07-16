package arrays

type DynamicArray[T any] struct {
	lentgh   int
	capacity int
	elemets  []T
}

func NewDynamicArray[T any](capacity int) *DynamicArray[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &DynamicArray[T]{
		lentgh:   0,
		capacity: capacity,
		elemets:  make([]T, capacity),
	}
}

func (a *DynamicArray[T]) Len() int {
	return a.lentgh
}

func (a *DynamicArray[T]) Capacity() int {
	return a.capacity
}

func (a *DynamicArray[T]) Push(value T) {
	if a.lentgh == a.capacity {
		a.capacity *= 2
		a.elemets = make([]T, a.capacity)
	}
	a.elemets[a.lentgh] = value
}

func (a *DynamicArray[T]) Insert(index int, value T) error {
	return nil
}
