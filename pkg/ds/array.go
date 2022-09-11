package ds

type list interface {}

type array[T list] struct {
	data []T
}

func NewArray[T list]() array[T] {
	var newList = make([]T, 0)
	return array[T]{
		data: newList,
	}
}

func (a *array[T]) Push(item T) {
	a.data = append(a.data, item)
}

func (a *array[T]) At(idx int) T {
	return a.data[idx]
}

func (a *array[T]) ForEach(callback func(item T, idx int)) {
	for idx, it := range a.data {
		callback(it, idx)
	}
}

func (a *array[T]) Map(callback func(item T, idx int) bool) array[T] {
	newArr := NewArray[T]()
	for idx, it := range a.data {
		if (callback(it, idx)) {
			newArr.Push(it)
		}
	}

	return newArr
}

func (a *array[T]) Pop() *T {
	if len(a.data) == 0 {
		return nil
	}

	popped := a.data[len(a.data) - 1]
	a.data = a.data[0:len(a.data) - 1]
	return &popped
}

func (a *array[T]) Shift() *T {
	if len(a.data) == 0 {
		return nil
	}

	shifted := a.data[0]
	a.data = a.data[1:len(a.data)]
	return &shifted
}

func (a *array[T]) Size() int {
	return len(a.data)
}

func (a *array[T]) Unshift(item T) {
	aux := a.data
	a.data = []T{item}
	a.data = append(a.data, aux...)
}
