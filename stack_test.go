package gods_test

import (
	"testing"
	"github.com/ParkerGits/gods"
)

func TestStack(t *testing.T) {
	stack := gods.NewStack[int]()

	if stack.Len() != 0 {
		t.Errorf("Length of empty stack should be 0")
	}

	stack.Push(4)
	if stack.Len() != 1 {
		t.Errorf("Length of stack after single push should be 1")
	}
	if stack.Peek() != 4 {
		t.Errorf("Top element after single push of 4 should be 4")
	}

	stack.Push(3)
	if stack.Len() != 2 {
		t.Errorf("Length of stack after two pushes should be 2")
	}
	if stack.Peek() != 3 {
		t.Errorf("Top element after second push should be 3")
	}

	popped := stack.Pop()
	if popped != 3 {
		t.Errorf("First element popped should be last element pushed (3)")
	}
	if stack.Len() != 1 {
		t.Errorf("Stack length should be 1 after two pushes and a pop")
	}
	if stack.Peek() != 4 {
		t.Errorf("Top of stack should be 4 after pop")
	}
}