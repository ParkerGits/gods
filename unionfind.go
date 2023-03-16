package gods

type UnionFind struct {
	rank []int
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{rank: make([]int, size), parent: parent}
}

func (uf *UnionFind) Find(element int) int {
	if element < 0 || element >= len(uf.parent) {
		panic("element out of range")
	}

	// path compression
	if element != uf.parent[element] {
		uf.parent[element] = uf.Find(uf.parent[element])
	}

	return uf.parent[element]
}


func (uf *UnionFind) Union(a int, b int) {
	nameA := uf.Find(a)
	nameB := uf.Find(b)
	if nameA == nameB {
		return;
	}

	if uf.rank[nameA] > uf.rank[nameB] {
		uf.parent[nameB] = nameA
	} else if uf.rank[nameA] < uf.rank[nameB] {
		uf.parent[nameA] = nameB
	} else {
		uf.parent[nameA] = uf.parent[nameB]
		uf.rank[nameB]++
	}
}