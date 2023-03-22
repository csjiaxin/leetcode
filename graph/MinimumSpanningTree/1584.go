package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Element is an entry in the priority queue
type E struct {
	x int
	y int
	v int
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

var visit []bool
var pq IntHeap
var adj [][]int
var N int

func add_edge(id int) {
	visit[id] = true
	for j := 0; j < N; j++ {
		if !visit[j] {
			e := E{id, j, adj[id][j]}
			heap.Push(&pq, e)
		}
	}
}
func minCostConnectPoints(points [][]int) int {
	N = len(points)
	visit = make([]bool, N)
	pq = make(IntHeap, 0)
	adj = make([][]int, N)
	for i := 0; i < N; i++ {
		adj[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			dis := int(math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1])))
			adj[i][j] = dis
			adj[j][i] = dis
		}
	}
	minCost := 0
	add_edge(0)
	for len(pq) > 0 {
		e := heap.Pop(&pq).(E)
		if !visit[e.y] {
			add_edge(e.y)
			minCost += e.v
		}
	}
	return minCost
}

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}
