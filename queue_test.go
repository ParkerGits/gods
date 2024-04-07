package gods_test

import (
	"github.com/ParkerGits/gods"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := gods.NewQueue[int]()

	if queue.Len() != 0 {
		t.Errorf("Length of empty queue should be 0")
	}

	queue.Enqueue(4)
	if queue.Len() != 1 {
		t.Errorf("Length of queue after single push should be 1")
	}
	if queue.Peek() != 4 {
		t.Errorf("Top element after single push of 4 should be 4")
	}

	queue.Enqueue(3)
	if queue.Len() != 2 {
		t.Errorf("Length of queue after two pushes should be 2")
	}
	if queue.Peek() != 4 {
		t.Errorf("Top element after second push should still be 4")
	}

	popped := queue.Dequeue()
	if popped != 4 {
		t.Errorf("First element popped should be first element pushed (4)")
	}
	if queue.Len() != 1 {
		t.Errorf("Queue length should be 1 after two pushes and a pop")
	}
	if queue.Peek() != 3 {
		t.Errorf("Top of queue should be 3 after pop")
	}
}
