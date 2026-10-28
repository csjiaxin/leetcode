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