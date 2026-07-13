package arrays

type Array struct {
	lentgh   uint64
	capacity uint64
	elemets  []int
}

func NewArray(size uint64) *Array {
	return &Array{
		lentgh:   0,
		capacity: size,
		elemets:  make([]int, size),
	}
}

func (a *Array) Insert(index uint64, value int) error {
	return nil
}
