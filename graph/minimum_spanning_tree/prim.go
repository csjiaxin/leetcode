package main

import (
	"container/heap"
	"fmt"
)

type edgeHeap []MSTEdge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *edgeHeap) Push(x any) {
	*h = append(*h, x.(MSTEdge))
}

func (h *edgeHeap) Pop() any {
	old := *h
	last := len(old) - 1
	edge := old[last]
	*h = old[:last]
	return edge
}

// 节点编号为 0 到 n-1，每条边视为无向边。
func Prim(n int, edges []MSTEdge) (int64, []MSTEdge, error) {
	if n <= 0 {
		return 0, nil, fmt.Errorf("n must be positive")
	}

	// 无向图：每条边加入两个方向。
	graph := make([][]MSTEdge, n)
	for _, edge := range edges {
		graph[edge.U] = append(graph[edge.U], edge)
		graph[edge.V] = append(graph[edge.V], MSTEdge{
			U:      edge.V,
			V:      edge.U,
			Weight: edge.Weight,
		})
	}

	visited := make([]bool, n)
	mst := make([]MSTEdge, 0, n-1)
	var total int64
	count := 0

	// 虚拟边：从节点 0 开始，不计入 MST。
	pq := &edgeHeap{{U: -1, V: 0, Weight: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(MSTEdge)
		u := edge.V

		if visited[u] {
			continue // 目标已经在树内，这条候选边已失效。
		}

		visited[u] = true
		count++

		if edge.U != -1 {
			next, ok := checkedAdd(total, edge.Weight)
			if !ok {
				return 0, nil, fmt.Errorf("total weight overflow")
			}
			total = next
			mst = append(mst, edge)
		}

		if count == n {
			return total, mst, nil
		}

		// 将新节点通向树外的边加入候选堆。
		for _, next := range graph[u] {
			if !visited[next.V] {
				heap.Push(pq, next)
			}
		}
	}

	return 0, nil, fmt.Errorf("graph is disconnected")
}