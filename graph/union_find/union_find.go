package unionfind

type UnionFind struct {
	parent []int
	size   []int
}

// Nodes are numbered 0 through n-1.
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.parent[x] {
		// Path compression: point x directly to its grandparent.
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	// Attach the smaller tree under the larger tree.
	if uf.size[rootA] < uf.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	uf.parent[rootB] = rootA
	uf.size[rootA] += uf.size[rootB]
	return true
}

func (uf *UnionFind) Connected(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}