package gods_test

import (
	"testing"
	"github.com/ParkerGits/gods"
)

func TestUnionFind(t *testing.T) {
	uf := gods.NewUnionFind(6)
	if uf.Find(2) != 2 {
		t.Errorf("all elements should initially be parents of themselves")
	}
	uf.Union(0, 1)
	uf.Union(0, 3)
	if uf.Find(1) != uf.Find(3) {
		t.Errorf("expected 1 and 3 to be part of the same set")
	}

	uf.Union(2, 4)
	uf.Union(2, 5)
	uf.Union(2, 0)
	if uf.Find(2) != uf.Find(1) {
		t.Errorf("expected 1 and 2 to be part of same set")
	}
}