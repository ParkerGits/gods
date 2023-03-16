package gods

var exists = struct{}{}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	set := Set[T]{elements: make(map[T]struct{})}
	return &set
}

func (set *Set[T]) Elements() map[T]struct{} {
	return set.elements
}

func (set *Set[T]) Contains(element T) bool {
	_, contains := set.elements[element]
	return contains
}

func (set *Set[T]) Add(element T) {
	set.elements[element] = exists
}

func (set *Set[T]) Len() int {
	return len(set.elements)
}

func (set *Set[T]) Remove(element T) {
	delete(set.elements, element)
}

func (set *Set[T]) ForEach(f func(element T)) {
	for element := range set.elements {
		f(element)
	}
}

func (set *Set[T]) From(arr []T) *Set[T] {
	for _, item := range arr {
		set.elements[item] = exists
	}
	return set
}