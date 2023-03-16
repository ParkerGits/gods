package gods

type BinaryHeapElement[K comparable] struct {
	Key int
	Value K
}

type BinaryHeap[K comparable] struct {
	elements []BinaryHeapElement[K]
	numElements int
	position map[K]int
}

func NewBinaryHeap[K comparable](capacity int) *BinaryHeap[K] {
	return &BinaryHeap[K]{elements: make([]BinaryHeapElement[K], capacity), numElements: 0, position: make(map[K]int)}
}

func (h *BinaryHeap[K]) Initialize(capacity int) {
	h.elements = make([]BinaryHeapElement[K], capacity)
	h.numElements = 0
	h.position = make(map[K]int)
}

func (h *BinaryHeap[K]) Len() int {
	return h.numElements
}

func (h *BinaryHeap[K]) ExtractMin() BinaryHeapElement[K] {
	return h.deleteIndex(0)
}

func (h *BinaryHeap[K]) Peek() BinaryHeapElement[K] {
	return h.elements[0]
}

func (h *BinaryHeap[K]) Insert(value K, key int) {
	if h.numElements == len(h.elements) {
		panic("heap full")
	}
	h.elements[h.numElements] = BinaryHeapElement[K]{key, value}
	h.position[value] = h.numElements
	h.numElements++
	h.heapifyUp(h.numElements-1)
}

func (h *BinaryHeap[K]) ChangeKey(value K, newKey int) bool {
	elementIndex, isInHeap := h.position[value]
	if !isInHeap {
		return false
	}
	oldKey := h.elements[elementIndex].Key
	if oldKey == newKey {
		return true
	}
	h.elements[elementIndex].Key = newKey
	if oldKey < newKey {
		h.heapifyDown(elementIndex)
	} else {
		h.heapifyUp(elementIndex)
	}
	return true
}

func (h *BinaryHeap[K]) DeleteElement(value K) (BinaryHeapElement[K], bool) {
	elementIndex, isInHeap := h.position[value]
	if !isInHeap {
		return BinaryHeapElement[K]{}, false
	}
	return h.deleteIndex(elementIndex), true
}

func (h *BinaryHeap[K]) KeyOf(value K) (int, bool) {
	elementIndex, isInHeap := h.position[value]
	if !isInHeap {
		return 0, false
	}
	return h.elements[elementIndex].Key, true
}

func (h *BinaryHeap[K]) deleteIndex(i int) BinaryHeapElement[K] {
	toRemove := h.elements[i]
	toSwap := h.elements[h.numElements-1]
	h.elements[i] = toSwap
	h.position[toSwap.Value] = i
	h.elements[h.numElements-1] = BinaryHeapElement[K]{}
	delete(h.position, toRemove.Value)
	h.numElements--
	h.heapifyDown(i)
	return toRemove
}

func (h* BinaryHeap[K]) leftChildIndex(i int) int {
	return 2*i+1
}

func (h* BinaryHeap[K]) parentIndex(i int) int {
	if i == 0 {
		return 0
	}
	return (i-1)/2
}

func (h *BinaryHeap[K]) heapifyUp(i int) {
	if i > 0 {
		pIndex := h.parentIndex(i)
		currElement := h.elements[i]
		parentElement := h.elements[pIndex]
		if currElement.Key < parentElement.Key {
			h.elements[pIndex] = currElement
			h.position[currElement.Value] = pIndex
			h.elements[i] = parentElement
			h.position[parentElement.Value] = i
			h.heapifyUp(pIndex)
		}
	}
}

func (h *BinaryHeap[K]) heapifyDown(i int) {
	lcIndex := h.leftChildIndex(i)
	if lcIndex > h.numElements - 1 {
		return
	}
	var indexToSwap int
	if lcIndex == h.numElements - 1 {
		indexToSwap = lcIndex
	} else {
		if h.elements[lcIndex].Key < h.elements[lcIndex + 1].Key {
			indexToSwap = lcIndex
		} else {
			indexToSwap = lcIndex + 1
		}
	}
	childElement := h.elements[indexToSwap]
	currElement := h.elements[i]
	if childElement.Key < currElement.Key {
		h.elements[indexToSwap] = currElement
		h.position[currElement.Value] = indexToSwap
		h.elements[i] = childElement
		h.position[childElement.Value] = i
		h.heapifyDown(indexToSwap)
	}
}