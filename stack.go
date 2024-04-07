package gods

type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

type Stack[T any] struct {
	length int
	head   *StackNode[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{length: 0, head: nil}
}

func (stack *Stack[T]) Push(value T) {
	new := &StackNode[T]{value: value, next: nil}
	if stack.length != 0 {
		new.next = stack.head
	}
	stack.head = new
	stack.length++
}

func (stack *Stack[T]) Pop() T {
	if stack.length == 0 {
		var zeroValue T
		return zeroValue
	}
	popped := stack.head
	stack.head = popped.next
	popped.next = nil
	stack.length--
	return popped.value
}

func (stack *Stack[T]) Peek() T {
	if stack.length == 0 {
		var zeroValue T
		return zeroValue
	}
	return stack.head.value
}

func (stack *Stack[T]) Len() int {
	return stack.length
}

func (stack *Stack[T]) Elements() []T {
	elements := make([]T, stack.length)
	i := 0
	for tmp := stack.head; tmp != nil; tmp = tmp.next {
		elements[i] = tmp.value
		i++
	}
	return elements
}
