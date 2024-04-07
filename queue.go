package gods

type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
	prev  *QueueNode[T]
}

type Queue[T any] struct {
	length int
	head   *QueueNode[T]
	tail   *QueueNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{length: 0, head: nil, tail: nil}
}

func (queue *Queue[T]) Enqueue(value T) {
	new := &QueueNode[T]{value: value, next: nil, prev: nil}
	if queue.length == 0 {
		queue.head = new
		queue.tail = new
		queue.length++
		return
	}

	queue.tail.next = new
	new.prev = queue.tail
	queue.tail = new
	queue.length++
}

func (queue *Queue[T]) Dequeue() T {
	if queue.length == 0 {
		var zeroValue T
		return zeroValue
	}

	popped := queue.head
	if queue.length == 1 {
		popped.next = nil
		queue.head = nil
		queue.tail = nil
		queue.length--
		return popped.value
	}

	popped.next.prev = nil
	queue.head = popped.next
	popped.next = nil
	queue.length--

	return popped.value
}

func (queue *Queue[T]) Peek() T {
	if queue.head == nil {
		var zeroValue T
		return zeroValue
	}
	return queue.head.value
}

func (queue *Queue[T]) Len() int {
	return queue.length
}

func (queue *Queue[T]) Elements() []T {
	elements := make([]T, queue.length)
	i := 0
	for tmp := queue.head; tmp != nil; tmp = tmp.next {
		elements[i] = tmp.value
		i++
	}
	return elements
}
