package main

import (
	"fmt"
	"sort"
)

type MSTEdge struct {
	U, V   int
	Weight int64
}

// 节点编号为 0 到 n-1，每条边视为无向边。
func Kruskal(n int, edges []MSTEdge) (int64, []MSTEdge, error) {
	if n <= 0 {
		return 0, nil, fmt.Errorf("n must be positive")
	}

	// 复制后排序，避免修改传入的 edges。
	sorted := append([]MSTEdge(nil), edges...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight < sorted[j].Weight
	})

	uf := NewUnionFind(n)
	mst := make([]MSTEdge, 0, n-1)
	var total int64

	for _, edge := range sorted {
		if !uf.Union(edge.U, edge.V) {
			continue // 两端已经连通，加入这条边会形成环。
		}

		next, ok := checkedAdd(total, edge.Weight)
		if !ok {
			return 0, nil, fmt.Errorf("total weight overflow")
		}

		total = next
		mst = append(mst, edge)

		if len(mst) == n-1 {
			break
		}
	}

	if len(mst) != n-1 {
		return 0, nil, fmt.Errorf("graph is disconnected")
	}

	return total, mst, nil
}

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

func checkedAdd(a, b int64) (int64, bool) {
	const highest = int64(1<<63 - 1)
	const lowest = int64(-1 << 63)
	if b > 0 && a > highest-b || b < 0 && a < lowest-b {
		return 0, false
	}
	return a + b, true
}
