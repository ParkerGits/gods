package gods_test

import (
	"testing"
	"github.com/ParkerGits/gods"
)

func TestBinaryHeap(t *testing.T) {
	heap := gods.NewBinaryHeap[int](10)
	heap.Insert(1, 6)
	if heap.Peek().Value != 1 {
		t.Errorf("Expected root element to have value 1. Actual: %d", heap.Peek().Value)
	}
	if heap.Len() != 1 {
		t.Errorf("Expected heap to have length 1. Actual: %d", heap.Len())
	}

	heap.Insert(2, 3)
	if heap.Peek().Value != 2 {
		t.Errorf("Expected root element to have value 2. Actual: %d", heap.Peek().Value)
	}
	if heap.Len() != 2 {
		t.Errorf("Expected heap to have length 2. Actual: %d", heap.Len())
	}
	heap.Insert(3, 4)
	if heap.Peek().Value != 2 {
		t.Errorf("Expected root element to have value 2. Actual: %d", heap.Peek().Value)
	}
	if heap.Len() != 3 {
		t.Errorf("Expected heap to have length 3. Actual: %d", heap.Len())
	}

	heap.Insert(4, 5)
	if heap.Peek().Value != 2 {
		t.Errorf("Expected root element to have value 2. Actual: %d", heap.Peek().Value)
	}
	if heap.Len() != 4 {
		t.Errorf("Expected heap to have length 4. Actual: %d", heap.Len())
	}

	heap.ChangeKey(4, 1)
	if heap.Peek().Value != 4 {
		t.Errorf("Expected root element to have value 4. Actual: %d", heap.Peek().Value)
	}

	min := heap.ExtractMin()
	if min.Value != 4 {
		t.Errorf("Expected root element to have value 4. Actual: %d", heap.Peek().Value)
	}
	if heap.Len() != 3 {
		t.Errorf("Expected heap to have length 3. Actual: %d", heap.Len())
	}

	heap.Insert(5, 2)
	heap.Insert(6, 7)
	heap.DeleteElement(2)
}