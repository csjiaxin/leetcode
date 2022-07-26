package main

import (
	"container/heap"
	"fmt"
)

type E struct {
	v int
	i int
}

// An IntHeap is a min-heap of ints.
type IntHeap []E

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(E))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumImportance(n int, roads [][]int) int64 {

	adj := make([]int, n)
	for _, e := range roads {
		adj[e[0]] += 1
		adj[e[1]] += 1
	}
	h := make(IntHeap, 0)
	heap.Init(&h)
	for i, v := range adj {
		heap.Push(&h, E{v, i})
	}

	id := 1
	ids := make([]int, n)
	for h.Len() > 0 {
		e := heap.Pop(&h).(E)
		ids[e.i] = id
		id += 1
	}

	r := int64(0)
	for _, e := range roads {
		r = r + int64(ids[e[0]]+ids[e[1]])
	}

	return r
}

func main() {
	fmt.Println(maximumImportance(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}))
}
