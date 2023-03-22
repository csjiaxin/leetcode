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

var visit [][]bool
var dir4 = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func valid(x, y, M, N int) bool {
	if x >= 0 && x < M && y >= 0 && y < N {
		return true
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumEffortPath(heights [][]int) int {
	queue := IntHeap([]E{}) // empty
	heap.Push(&queue, E{0, 0, 0})
	dist := make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dist[i] = make([]int, len(heights[0]))
	}
	M := len(heights)
	N := len(heights[0])
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			dist[i][j] = math.MaxInt
		}
	}
	dist[0][0] = 0
	visit = make([][]bool, M)
	for i := 0; i < M; i++ {
		visit[i] = make([]bool, N)
	}

	for len(queue) > 0 {
		de := heap.Pop(&queue)
		e := de.(E)
		if visit[e.x][e.y] {
			continue
		}
		visit[e.x][e.y] = true
		dist[e.x][e.y] = e.v
		for _, d := range dir4 {
			x := e.x + d[0]
			y := e.y + d[1]

			if valid(x, y, M, N) && visit[x][y] == false {
				effort := int(math.Abs(float64(heights[x][y] - heights[e.x][e.y])))
				heap.Push(&queue, E{x, y, max(e.v, effort)})
			}
		}
	}
	return dist[M-1][N-1]

}

func main() {
	// fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
}
