package gods_test

import (
	"testing"
	"github.com/ParkerGits/gods"
)

func TestSet(t *testing.T) {
	set := gods.NewSet[int]()
	if set.Len() != 0 {
		t.Errorf("Length of empty set should be zero.")	
	}
	set.Add(1)
	if set.Len() != 1 {
		t.Errorf("Length of set with one element should be one.")
	}
	if !set.Contains(1) {
		t.Errorf("Set should contain the element 1")
	}

	set.Add(2)
	set.Add(3)
	if set.Len() != 3 {
		t.Errorf("Length of set with three elements should be three.")
	}
	if !set.Contains(2) || !set.Contains(3) {
		t.Errorf("Set should contain elements 2 and 3")
	}

	set.Remove(2)
	if set.Len() != 2 {
		t.Errorf("Length of set with two elements should be two.")
	}
	if set.Contains(2) {
		t.Errorf("Set should not contain 2 after remove")
	}
}